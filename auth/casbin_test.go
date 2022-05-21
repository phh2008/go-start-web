package auth

import (
	"com.phh/start-web/pkg/config"
	"com.phh/start-web/util"
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

var enforcer *casbin.Enforcer

func init() {
	config := config.NewConfig(".././config")
	db := util.InitDB(config)
	enforcer = NewCasbinHelper(db, "../config").GetEnforcer()
}

func Test1(t *testing.T) {
	m := enforcer.GetModel()
	fmt.Println(m)
}

// 增加策略
func Test2(t *testing.T) {
	if ok, _ := enforcer.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
		fmt.Println("Policy已经存在")
	} else {
		fmt.Println("增加成功")
	}
}

// 删除策略
func Test3(t *testing.T) {
	fmt.Println("删除Policy")
	if ok, _ := enforcer.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
		fmt.Println("Policy不存在")
	} else {
		fmt.Println("删除成功")
	}
}

// 获取策略
func Test4(t *testing.T) {
	fmt.Println("查看policy")
	list := enforcer.GetPolicy()
	for _, vlist := range list {
		for _, v := range vlist {
			fmt.Printf("value: %s, ", v)
		}
		fmt.Printf("\n")
	}
}

// 更新
func Test5(t *testing.T) {
	enforcer.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data1", "write"})
}

// 查询，条件过滤
func Test6(t *testing.T) {
	// 更新策略，根据条件(fieldIndex：表示参数开始字段索引)
	// 生成 sql如下，
	// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'bob' and v1 = 'data2'
	// INSERT INTO `casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`) VALUES ('p','admin','/api/v3/hello','POST','','','','','')
	enforcer.UpdateFilteredPolicies([][]string{{"admin", "/api/v3/hello", "POST"}}, 0, "bob", "data2")
}
