package model

import "gopkg.in/mgo.v2/bson"

/********************************* 游戏角色属性 ************************************/
type (
	// 游戏角色
	Player struct {
		Id            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Nickname      string        `bson:"nickname,omitempty" json:"nickname"`   // 昵称
		Level         int           `bson:"level" json:"level"`                   // 等级
		Experience    int64         `bson:"experience" json:"experience"`         // 本等级经验
		Money         int           `bson:"money" json:"money"`                   // 金钱
		Bag           int           `bson:"bag" json:"bag"`                       // 背包数量
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

		MsgChan chan string `bson:"msg" json:"msg"`
	}

	// 角色五行属性
	Element struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Uid   string `bson:"pid,omitempty" json:"pid"` // 角色ID
		Metal int           `bson:"metal" json:"metal"`       // 金
		Wood  int           `bson:"wood" json:"wood"`         // 木
		Water int           `bson:"water" json:"water"`       // 水
		Fire  int           `bson:"fire" json:"fire"`         // 火
		Earth int           `bson:"fire" json:"fire"`         // 土
		Sort  int           `bson:"sort" json:"sort"` 				// 排序
	}

	// 角色技能
	Skill struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Rid   string `bson:"rid,omitempty" json:"rid"` // 角色ID
		Sid   string `bson:"sid,omitempty" json:"sid"` // 技能ID
		Level int           `bson:"level" json:"level"`       // 技能等级
		Sort  int           `bson:"sort" json:"sort"` 				// 排序
	}
)

/***************************************** 物品 *********************************/
type (
	// 物品
	Item struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Cid      string `bson:"cid,omitempty" json:"cid"`   // 类型ID
		Name     string        `bson:"name,omitempty" json:"name"` // 物品名称
		Level	 int           `bson:"level" json:"level"`		   // 物品等级
		Function string        `bson:"function" json:"function"`   // 物品功能
		Price    int           `bson:"price" json:"price"`         // 物品基本卖价
		FallNum  int64         `bson:"fall_num" json:"fall_num"`   // 掉落次数
		Abstract string        `bson:"abstract" json:"abstract"`   // 物品简介
	}

	// 类型
	ItemCate struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Pid   string `bson:"pid,omitempty" json:"pid"`				// 父ID
		Name  string        `bson:"name,omitempty" json:"name"`
		Sort  int           `bson:"sort" json:"sort"` 				// 排序
		Abstract string     `bson:"abstract" json:"abstract"`   	// 简介
	}
)

/**************************************** 装备 *********************************/
