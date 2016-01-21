package model

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)


func (m *ItemCate) mgo() *Mgo {
	return NewMgo("items")
}

// 新增物品分类
func (m *ItemCate) AddItemCate() error {
	c := m.mgo()
	if !m.CheckName() {
		return fmt.Errorf("item cate name 不能重复!")
	}
	return c.Insert(m)
}

// 检测分类名称
func (m *ItemCate) CheckName() bool {
	err := m.FindByName()
	if err != nil {
		return true
	}

	return false
}

// 根据名称查询分类
func (m *ItemCate) FindByName() error {
	c := m.mgo()
	c.Query = bson.M{"name": m.Name}
	return c.Find(&m)
}

// 根据ID查询分类
func (m *ItemCate) FindById() error {
	c := m.mgo()
	c.Query = bson.M{"_id": m.Id}
	return c.Find(&m)
}

// 查询所有物品信息
// @param sort 排序
// @param sel select 返回指定字段，如{"name":1} 只返回name字段 {"name":0} 不返回name字段
// @param limit 返回文档个数
// @param skip 跳过文档个数
func (m *ItemCate) FindAll(skip, limit int, sort ...string) (v []*ItemCate, err error) {
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

// 根据名称更新分类
func (m *ItemCate) UpdateByName(name string) error {
	c := m.mgo()
	c.Query = bson.M{"name": name}
	c.Change = bson.M{"$set": m}

	return c.Update()
}

// 根据ID更新分类
func (m *ItemCate) UpdateById(id string) error {
	c := m.mgo()
	c.Query = bson.M{"_id": bson.IsObjectIdHex(id)}
	c.Change = bson.M{"$set": m}

	return c.Update()
}

// 根据名称删除分类
func (m *ItemCate) DelByName(name string) error {
	c := m.mgo()
	c.Query = bson.M{"name": name}

	return c.Remove()
}

// 根据ID删除分类
func (m *ItemCate) DelById(id string) error {
	c := m.mgo()
	c.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return c.Remove()
}