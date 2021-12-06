package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	dc            *mongo.Collection
	ac            *mongo.Collection
	enableCache   bool
	notifyCachPre string
	counter       counter.Service
	user          user.UserServiceServer
	role          role.RoleServiceServer
	log           logger.Logger

	department.UnimplementedDepartmentServiceServer
}

func (s *service) Config() error {
	if pkg.Counter == nil {
		return fmt.Errorf("dependence counter service is nil")
	}
	s.counter = pkg.Counter
	if pkg.User == nil {
		return fmt.Errorf("dependence user service is nil")
	}
	s.user = pkg.User

	if pkg.Role == nil {
		return fmt.Errorf("dependence role service is nil")
	}
	s.role = pkg.Role

	db := conf.C().Mongo.GetDB()

	dc := db.Collection("department")
	dcIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "domain", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err := dc.Indexes().CreateMany(context.Background(), dcIndexs)
	if err != nil {
		return err
	}

	ac := db.Collection("join_apply")
	acIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = ac.Indexes().CreateMany(context.Background(), acIndexs)
	if err != nil {
		return err
	}

	s.dc = dc
	s.ac = ac
	s.log = zap.L().Named("Department")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return department.HttpEntry()
}

func init() {
	pkg.RegistryService("department", Service)
}
