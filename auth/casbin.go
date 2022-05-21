package auth

import (
	"com.phh/start-web/pkg/config"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"
	"gorm.io/gorm"
	"log"
	"path/filepath"
	"sync"
)

var CasbinSet = wire.NewSet(wire.Struct(new(CasbinHelper), "*"))

var rbacEnforcer *casbin.Enforcer
var once sync.Once

type CasbinHelper struct {
	Db      *gorm.DB
	ConfDir config.ConfigFolder
}

func NewCasbinHelper(db *gorm.DB, config config.ConfigFolder) *CasbinHelper {
	return &CasbinHelper{Db: db, ConfDir: config}
}

func (a *CasbinHelper) GetEnforcer() *casbin.Enforcer {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(a.Db)
		if err != nil {
			log.Println("casbin gorm 适配器创建失败")
			panic(err)
		}
		configFile := filepath.Join(string(a.ConfDir), "rbac_model.conf")
		fmt.Println(">>>>modelFile:", configFile)
		rbacEnforcer, err = casbin.NewEnforcer(configFile, adapter)
		if err != nil {
			log.Println("casbin 创建失败")
			panic(err)
		}
		rbacEnforcer.EnableAutoSave(true)
		// Load the policy from DB.
		_ = rbacEnforcer.LoadPolicy()
	})
	return rbacEnforcer
}
