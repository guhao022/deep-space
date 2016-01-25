package model

import "gopkg.in/mgo.v2/bson"

/********************************* 游戏角色属性 ************************************/

// 游戏角色
type Player struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nickname   string        `bson:"nickname,omitempty" json:"nickname"` // 昵称
	Level      int           `bson:"level" json:"level"`                 // 等级
	Experience int64         `bson:"experience" json:"experience"`       // 经验
	Money      int           `bson:"money" json:"money"`                 // 金钱
	Bag        int         	 `bson:"bag" json:"bag"`                     // 背包数量
	MsgChan    chan string   `bson:"msg" json:"msg"`
}

// 角色技能
type PSkill struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Uid   bson.ObjectId `bson:"uid,omitempty" json:"uid"` // 角色ID
	Sid   bson.ObjectId `bson:"sid,omitempty" json:"sid"` // 技能ID
	Level int           `bson:"level" json:"level"`       // 技能等级
	Sort  int           `bson:"sort" json:"sort"`         // 排序
}

type PEquipment struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Uid       bson.ObjectId `bson:"uid,omitempty" json:"uid"`             // 角色ID
	Equipment Item     `bson:"equipment,omitempty" json:"equipment"` // 装备ID
}

// 角色身上的装备
type PAttr struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Uid   bson.ObjectId `bson:"uid,omitempty" json:"uid"` // 角色ID
	Attrid   bson.ObjectId `bson:"attr_id,omitempty" json:"attr_id"` // 装备ID
}

// 背包
type Bag struct {
	Id  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Uid   bson.ObjectId `bson:"uid,omitempty" json:"uid"` // 角色ID
	ItemId bson.ObjectId `bson:"item_id,omitempty" json:"item_id"`	// 物品ID

}

/***************************************** 属性 *********************************/

// 属性
type Attribute struct {
	Id            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Hp            int64         `bson:"hp" json:"hp"`                         // 血量
	Mp            int64         `bson:"mp" json:"mp"`                         // 魔力值(技能消耗)
	Agile         int           `bson:"agile" json:"agile"`                   // 敏捷
	Strengh       int           `bson:"strengh" json:"strengh"`               // 力量
	Intellect     int           `bson:"intellect" json:"intellect"`           // 智力
	Physique      int           `bson:"physique" json:"physique"`             // 体质
	PhyCounteract float64       `bson:"phy_counteract" json:"phy_counteract"` // 物理抗性
	MagCounteract float64       `bson:"mag_counteract" json:"mag_counteract"` // 魔法抗性
	HitPoint      int           `bson:"hit_point" json:"hit_point"`           // 命中率
	FleePoint     int           `bson:"flee_point" json:"flee_point"`         // 回避率
	PhyMax        int           `bson:"phy_max" json:"phy_max"`               // 物理最大攻击
	PhyMin        int           `bson:"phy_min" json:"phy_min"`               // 物理最小攻击
	MagMax        int           `bson:"mag_max" json:"mag_max"`               // 魔法最大攻击
	MagMin        int           `bson:"mag_min" json:"mag_min"`               // 魔法最小攻击
	Element       Element       `json:"element"`                              // 五行
}

// 角色五行属性
type Element struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Metal int           `bson:"metal" json:"metal"` // 金
	Wood  int           `bson:"wood" json:"wood"`   // 木
	Water int           `bson:"water" json:"water"` // 水
	Fire  int           `bson:"fire" json:"fire"`   // 火
	Earth int           `bson:"fire" json:"fire"`   // 土
}



/***************************************** 技能 *********************************/

type Skill struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name,omitempty" json:"name"` // 技能名称
	Unlock  int `bson:"unlock" json:"unlock"` // 技能解锁等级
	MaxLevel  int `bson:"max_level" json:"max_level"` // 技能最高等级
}

/***************************************** 物品 *********************************/

type (
// 物品
	Item struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Cid      bson.ObjectId `bson:"cid,omitempty" json:"cid"`   // 类型ID
		Name     string        `bson:"name,omitempty" json:"name"` // 物品名称
		Attribute Attribute     `bson:"attribute" json:"attribute"`	// 属性
		Body      int           `bson:"body json:"body"`            // 装备位置
		Level	 int           `bson:"level" json:"level"`		   // 物品等级
		Price    int           `bson:"price" json:"price"`         // 物品基本卖价
		IsFall	 bool            `bson:"is_fall" json:"is_fall"`	// 是否掉落
		FallNum  int64         `bson:"fall_num" json:"fall_num"`   // 掉落次数
		Summary string     	   `bson:"summary" json:"summary"`   // 物品简介
	}

// 类型
	ItemCate struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `bson:"name,omitempty" json:"name"`
		Sort  int           `bson:"sort" json:"sort"` 				// 排序
		Summary string     `bson:"summary" json:"summary"`   	// 简介
	}
)
