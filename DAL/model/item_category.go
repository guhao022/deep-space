package model
import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var (
	itemCateC = Mgo.C("item_category")
)

// 新增物品分类
func (m *ItemCate) AddItemCate() error {
	if !m.CheckName() {
		return fmt.Errorf("item cate name 不能重复!")
	}
	return itemCateC.Insert(m)
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
	itemCateC.Query = bson.M{"name": m.Name}
	return itemCateC.Find(&m)
}

// 根据ID查询分类
func (m *ItemCate) FindById() error {
	itemCateC.Query = bson.M{"_id": m.Id}
	return itemCateC.Find(&m)
}

// 查询所有物品信息
// @param sort 排序
// @param sel select 返回指定字段，如{"name":1} 只返回name字段 {"name":0} 不返回name字段
// @param limit 返回文档个数
// @param skip 跳过文档个数
func (m *ItemCate) FindAll(skip, limit int, sort ...string) (v []*ItemCate, err error) {
	if len(sort) > 0 {
		itemCateC.Sort = sort
	}
	if skip > 0 {
		itemCateC.Skip = skip
	}
	if limit > 0 {
		itemCateC.Limit = limit
	}

	err = itemCateC.FindAll(&v)
	return v, err
}

// 根据名称更新分类
func (m *ItemCate) UpdateByName(name string) error {
	itemCateC.Query = bson.M{"name": name}
	itemCateC.Change = bson.M{"$set": m}

	return itemCateC.Update()
}

// 根据ID更新分类
func (m *ItemCate) UpdateById(id string) error {
	itemCateC.Query = bson.M{"_id": bson.IsObjectIdHex(id)}
	itemCateC.Change = bson.M{"$set": m}

	return itemCateC.Update()
}

// 根据名称删除分类
func (m *ItemCate) DelByName(name string) error {
	itemCateC.Query = bson.M{"name": name}

	return itemCateC.Remove()
}

// 根据ID删除分类
func (m *ItemCate) DelById(id string) error {
	itemCateC.Query = bson.M{"_id": bson.ObjectIdHex(id)}

	return itemCateC.Remove()
}