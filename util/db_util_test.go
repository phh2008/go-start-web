package util

import (
	"com.phh/start-web/config"
	"com.phh/start-web/entity"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func init() {
	var config = config.NewConfig("../config")
	db = InitDB(config)
}

// 子查询，string 条件
func TestQuery1(t *testing.T) {
	var user []map[string]interface{}
	db.Table("(?) tmp",
		db.Model(&entity.User{}).
			Select("id", "name").
			Where("age>?", 18),
	).Where("tmp.name=?", "tom").Find(&user)
	fmt.Println(user)
}

// struct、map条件
func TestQuery2(t *testing.T) {
	// struct
	var users entity.User
	db.Where(&entity.User{Name: "tom"}).Find(&users)
	fmt.Println(users)

	// map
	var users2 entity.User
	db.Where(map[string]interface{}{"name": "jack"}).Find(&users2)
	fmt.Println(users2)
}

// order 、group、having
func TestQuery3(t *testing.T) {
	var orders []map[string]interface{}
	db.Model(&entity.Order{}).
		Select("user_id", "sum(amount) as totalAmount", "count(user_id) as orderQty").
		Group("user_id").
		Having("orderQty>?", 2).
		Order("totalAmount desc").Find(&orders)
	fmt.Println(orders)
}

// join
func TestQuery4(t *testing.T) {
	var _ = `
SELECT
	a.id,
	a.user_id,
	a.amount,
	a.create_at,
	a.status,
	b.name 
FROM
	order a
	JOIN USER b ON a.user_id = b.id 
	AND a.status = 4 
WHERE
	a.user_id = 2 
ORDER BY
	a.amount DESC
	`
	var users []map[string]interface{}
	db.Table("`order` a").
		Select("a.id", "a.user_id", "a.amount", "a.create_at", "a.status", "b.name").
		Joins("join user b on a.user_id=b.id and a.status=?", 4).
		Where("a.user_id = ?", 2).
		Order("a.amount desc").Find(&users)
	fmt.Println(users)
}

// subquery + join
func TestQuery5(t *testing.T) {
	var users []entity.User
	// subquery
	subquery := db.Model(&entity.Order{}).Select("user_id").Group("user_id").Having("count(1)>?", 6)
	// join
	db.Table("user a").
		Select("a.id", "a.name", "a.salt", "a.age", "a.passwd", "a.birthday", "a.created").
		Joins("join (?) b ON a.id = b.user_id", subquery).Find(&users)
	fmt.Println(users)
}

// 更新
func TestUpdate1(t *testing.T) {
	//gdb.Model(&MemUser{}).Where("id>=?", 11111).Update("state", "2")
	// 表达式
	db.Model(&entity.User{}).Where("id>=?", 1).Update("age", gorm.Expr("age+?", 1))
}

// 根据条件和 model 的值进行更新
func TestUpdate2(t *testing.T) {
	// 根据条件和 model 的值进行更新
	var user = entity.User{Id: 1}
	db.Model(&user).Where("name = ?", "tom").Update("name", "hello")
	// UPDATE `user` SET `name`='hello' WHERE name = 'tom' AND `id` = 1
}

func TestUpdate3(t *testing.T) {
	// Select 除 passwd 外的所有字段（包括零值字段的所有字段）
	var user = entity.User{Id: 1}
	db.Model(&user).Select("*").Omit("passwd").Updates(entity.User{Name: "jinzhu", Passwd: "123", Age: 0})
	// UPDATE `user` SET `id`=0,`name`='jinzhu',`salt`='',`age`=0,`birthday`='0000-00-00 00:00:00',`created`='0000-00-00 00:00:00',`updated`='0000-00-00 00:00:00' WHERE `id` = 1
}
