package model
import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var (
	itemC = Mgo.C("items")
)

// 新增物品
func (m *Item) AddItem() error {
	if !m.CheckName() {
		return fmt.Errorf("item name 不能重复!")
	}
	return itemC.Insert(m)
}

// 检查物品名称是否重复
func (m *Item) CheckName() bool {
	err := m.FindByName()
	if err != nil {
		return true
	}

	return false
}

// 根据名称查询物品
func (m *Item) FindByName() error {
	itemC.Query = bson.M{"name": m.Name}
	return itemC.Find(&m)
}

// 根据ID查询物品
func (m *Item) FindByID() error {
	itemC.Query = bson.M{"_id": m.Id}
	return itemC.Find(&m)
}

// 根据分类查询物品
func (m *Item) FindByCID() error {
	itemC.Query = bson.M{"cid": m.Cid}
	return itemC.Find(&m)
}

// 查询所有物品信息
// @param sort 排序
// @param sel select 返回指定字段，如{"name":1} 只返回name字段 {"name":0} 不返回name字段
// @param limit 返回文档个数
// @param skip 跳过文档个数
func (m *Item) FindAll(skip, limit int, sel interface{}, sort ...string) (v []*Item, err error) {
	if len(sort) > 0 {
		itemC.Sort = sort
	}
	if skip > 0 {
		itemC.Skip = skip
	}
	if limit > 0 {
		itemC.Limit = limit
	}

	err = itemC.FindAll(&v)
	return v, err
}

// 根据名称更改物品信息
func (m *Item) UpdateByName(name string) error {
	itemC.Query = bson.M{"name": name}
	itemC.Change = bson.M{"$set": m}

	return itemC.Update()
}

// 根据ID更改物品信息
func (m *Item) UpdateById(id string) error {
	itemC.Query = bson.M{"_id": bson.IsObjectIdHex(id)}
	itemC.Change = bson.M{"$set": m}

	return itemC.Update()
}

// 根据名称删除物品
func (m *Item) DelByName(name string) error {
	itemC.Query = bson.M{"name": name}

	return itemC.Remove()
}

// 根据ID删除数据
func (m *Item) DelById(id string) error {
	itemC.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return itemC.Remove()
}