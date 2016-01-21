package model

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func (m *Item) mgo() *Mgo {
	return NewMgo("items")
}

// 新增物品
func (m *Item) AddItem() error {
	c := m.mgo()
	fmt.Println("c:"+c.Collection)
	if !m.CheckName() {
		return fmt.Errorf("item name 不能重复!")
	}
	return c.Insert(m)
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
	c := m.mgo()
	c.Query = bson.M{"name": m.Name}
	return c.Find(&m)
}

// 根据ID查询物品
func (m *Item) FindByID() error {
	c := m.mgo()
	c.Query = bson.M{"_id": m.Id}
	return c.Find(&m)
}

// 根据分类查询物品
func (m *Item) FindByCID() error {
	c := m.mgo()
	c.Query = bson.M{"cid": m.Cid}
	return c.Find(&m)
}

// 查询所有物品信息
// @param sort 排序
// @param sel select 返回指定字段，如{"name":1} 只返回name字段 {"name":0} 不返回name字段
// @param limit 返回文档个数
// @param skip 跳过文档个数
func (m *Item) FindAll(skip, limit int, sel interface{}, sort ...string) (v []*Item, err error) {
	c := m.mgo()
	if len(sort) > 0 {
		c.Sort = sort
	}
	if skip > 0 {
		c.Skip = skip
	}
	if limit > 0 {
		c.Limit = limit
	}

	err = c.FindAll(&v)
	return v, err
}

// 根据名称更改物品信息
func (m *Item) UpdateByName(name string) error {
	c := m.mgo()
	c.Query = bson.M{"name": name}
	c.Change = bson.M{"$set": m}

	return c.Update()
}

// 根据ID更改物品信息
func (m *Item) UpdateById(id string) error {
	c := m.mgo()
	c.Query = bson.M{"_id": bson.IsObjectIdHex(id)}
	c.Change = bson.M{"$set": m}

	return c.Update()
}

// 根据名称删除物品
func (m *Item) DelByName(name string) error {
	c := m.mgo()
	c.Query = bson.M{"name": name}

	return c.Remove()
}

// 根据ID删除数据
func (m *Item) DelById(id string) error {
	c := m.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return c.Remove()
}