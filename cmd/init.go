package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcube/http/label"
	"github.com/spf13/cobra"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/system"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

// InitCmd 初始化系统
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化服务",
	Long:  `初始化admin用户相关基础信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogzap(); err != nil {
			return err
		}

		// 加载缓存
		if err := loadCache(); err != nil {
			return err
		}

		// 初始化服务层
		if err := pkg.InitService(); err != nil {
			return err
		}

		initer, err := NewInitialerFromCLI()
		if err != nil {
			return err
		}
		if err := initer.Run(); err != nil {
			return err
		}
		return nil
	},
}

// NewInitialerFromCLI 初始化
func NewInitialerFromCLI() (*Initialer, error) {
	i := NewInitialer()

	if err := i.checkIsInit(); err != nil {
		return nil, err
	}

	err := survey.AskOne(
		&survey.Input{
			Message: "请输入公司(组织)名称:",
			Default: "USTB",
		},
		&i.domainDisplayName,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入管理员用户名称:",
			Default: "admin",
		},
		&i.username,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	var repeatPass string
	err = survey.AskOne(
		&survey.Password{
			Message: "请输入管理员密码:",
		},
		&i.password,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}
	err = survey.AskOne(
		&survey.Password{
			Message: "再次输入管理员密码:",
		},
		&repeatPass,
		survey.WithValidator(survey.Required),
		survey.WithValidator(func(ans interface{}) error {
			if ans.(string) != i.password {
				return fmt.Errorf("两次输入的密码不一致")
			}
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// NewInitialer todo
func NewInitialer() *Initialer {
	return &Initialer{
		mockTK: &token.Token{
			UserType: types.UserType_SUPPER,
			Domain:   domain.AdminDomainName,
		},
	}
}

// Initialer 初始化控制器
type Initialer struct {
	domainDisplayName string
	username          string
	password          string
	tk                *token.Token
	mockTK            *token.Token
}

func (i *Initialer) mockContext(account string) context.Context {
	ctx := pkg.NewGrpcInCtx()
	ctx.SetIsInternalCall(account, domain.AdminDomainName)
	return ctx.Context()
}

func (i *Initialer) userContext() context.Context {
	ctx := pkg.NewGrpcInCtx()
	ctx.SetAccessToken(i.tk.AccessToken)
	return ctx.Context()
}

// Run 执行初始化
func (i *Initialer) Run() error {
	fmt.Println()
	fmt.Println("开始初始化...")

	u, err := i.initUser()
	if err != nil {
		return err
	}
	fmt.Printf("初始化用户: %s [成功]\n", i.username)

	_, err = i.initDomain(u.Account)
	if err != nil {
		return err
	}
	fmt.Printf("初始化域: %s   [成功]\n", i.domainDisplayName)

	apps, err := i.initApp(u.Account)
	if err != nil {
		return err
	}
	for index := range apps {
		fmt.Printf("初始化应用: %s [成功]\n", apps[index].Name)
		fmt.Printf("应用客户端ID: %s\n", apps[index].ClientId)
		fmt.Printf("应用客户端凭证: %s\n", apps[index].ClientSecret)
	}

	if err := i.getAdminToken(apps[0]); err != nil {
		return err
	}

	var adminRole *role.Role
	roles, err := i.initRole()
	if err != nil {
		return err
	}
	for index := range roles {
		r := roles[index]
		fmt.Printf("初始化角色: %s [成功]\n", r.Name)
		if r.Name == role.AdminRoleName {
			adminRole = r
		}
	}

	svr, err := i.initService(adminRole)
	if err != nil {
		return err
	}
	fmt.Printf("初始化服务: %s   [成功]\n", svr.Name)

	dep, err := i.initDepartment()
	if err != nil {
		return err
	}
	fmt.Printf("初始化根部门: %s   [成功]\n", dep.DisplayName)

	sysconf, err := i.initSystemConfig()
	if err != nil {
		return err
	}
	fmt.Printf("初始化系统配置: %s   [成功]\n", sysconf.Version)

	return nil
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest()
	req.UserType = types.UserType_SUPPER
	userSet, err := pkg.User.QueryAccount(i.mockContext("internal"), req)
	if err != nil {
		return err
	}

	if userSet.Total > 0 {
		return errors.New("supper admin user has exist")
	}
	return nil
}

func (i *Initialer) initUser() (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.UserType = types.UserType_SUPPER
	req.Account = strings.TrimSpace(i.username)
	req.Password = strings.TrimSpace(i.password)
	return pkg.User.CreateAccount(i.mockContext(i.username), req)
}

func (i *Initialer) initDomain(account string) (*domain.Domain, error) {
	req := domain.NewCreateDomainRequest()
	req.Name = domain.AdminDomainName
	req.Profile.DisplayName = strings.TrimSpace(i.domainDisplayName)
	return pkg.Domain.CreateDomain(i.mockContext(account), req)
}

func (i *Initialer) initApp(account string) ([]*application.Application, error) {
	req := application.NewCreateApplicatonRequest()
	req.Name = application.AdminWebApplicationName
	req.Domain = "admin"
	req.ClientType = application.ClientType_PUBLIC
	req.Description = "Admin Web管理端"
	req.BuildIn = true

	ctx := i.mockContext(account)
	web, err := pkg.Application.CreateApplication(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("1. create admin web applicaton error, %s", err)
	}

	req = application.NewCreateApplicatonRequest()
	req.Name = application.AdminServiceApplicationName
	req.Domain = "admin"
	req.ClientType = application.ClientType_CONFIDENTIAL
	req.Description = "Admin Service 内置管理端, 服务注册后, 使用该端管理他们的凭证, 默认token不过期"
	req.AccessTokenExpireSecond = 0
	req.RefreshTokenExpireSecond = 0
	req.BuildIn = true
	svr, err := pkg.Application.CreateApplication(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("2. create admin web applicaton error, %s", err)
	}

	apps := []*application.Application{web, svr}
	return apps, nil
}

func (i *Initialer) getAdminToken(app *application.Application) error {
	if app == nil {
		return fmt.Errorf("get admin token need app")
	}

	req := token.NewIssueTokenByPassword(app.ClientId, app.ClientSecret, i.username, i.password)
	tk, err := pkg.Token.IssueToken(context.Background(), req)
	if err != nil {
		return err
	}
	i.tk = tk
	return nil
}

func (i *Initialer) initRole() ([]*role.Role, error) {
	admin := role.NewDefaultPermission()
	admin.ServiceId = "*"
	admin.ResourceName = "*"
	admin.LabelKey = "*"
	admin.LabelValues = []string{"*"}

	req := role.NewCreateRoleRequest()
	req.Name = role.AdminRoleName
	req.Description = "系统管理员, 有系统所有功能的访问权限"
	req.Permissions = []*role.CreatePermssionRequest{admin}
	req.Type = role.RoleType_BUILDIN
	adminRole, err := pkg.Role.CreateRole(i.userContext(), req)
	if err != nil {
		return nil, err
	}

	vistor := role.NewDefaultPermission()
	vistor.ServiceId = "*"
	vistor.ResourceName = "*"
	vistor.LabelKey = label.ActionLableKey
	vistor.LabelValues = []string{label.Get.Value(), label.List.Value()}

	req = role.NewCreateRoleRequest()
	req.Name = role.VisitorRoleName
	req.Description = "访客, 登录系统后, 默认的权限"
	req.Permissions = []*role.CreatePermssionRequest{vistor}
	req.Type = role.RoleType_BUILDIN
	vistorRole, err := pkg.Role.CreateRole(i.userContext(), req)
	if err != nil {
		return nil, err
	}

	return []*role.Role{adminRole, vistorRole}, nil
}

func (i *Initialer) initDepartment() (*department.Department, error) {
	if pkg.Department == nil {
		return nil, fmt.Errorf("dependence service department is nil")
	}

	req := department.NewCreateDepartmentRequest()
	req.Name = department.DefaultDepartmentName
	req.DisplayName = i.domainDisplayName
	req.Manager = strings.TrimSpace(i.username)
	return pkg.Department.CreateDepartment(i.userContext(), req)
}

func (i *Initialer) initService(r *role.Role) (*micro.Micro, error) {
	req := micro.NewCreateMicroRequest()
	req.Name = version.ServiceName
	req.Description = version.Description
	req.Type = micro.Type_BUILD_IN
	return pkg.Micro.CreateService(i.userContext(), req)
}

func (i *Initialer) initSystemConfig() (*system.Config, error) {
	sysConf := system.NewDefaultConfig()
	if err := pkg.System.InitConfig(sysConf); err != nil {
		return nil, err
	}
	return sysConf, nil
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
