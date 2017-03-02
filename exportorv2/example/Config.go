// Generated by github.com/davyxu/tabtoy
// Version: 2.7.2
// DO NOT EDIT!!
package table

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

// Defined in table: Globals
type ActorType int32

const (

	// 唐僧
	ActorType_Leader ActorType = 0

	// 孙悟空
	ActorType_Monkey ActorType = 1

	// 猪八戒
	ActorType_Pig ActorType = 2

	// 沙僧
	ActorType_Hammer ActorType = 3
)

// Defined in table: Config
type Config struct {

	//Sample
	Sample []*SampleDefine

	//Vertical
	Vertical []*VerticalDefine

	//Exp
	Exp []*ExpDefine
}

// Defined in table: Sample
type Prop struct {

	// 血量
	HP int32

	// 攻击速率
	AttackRate float32

	// 额外类型
	ExType ActorType
}

// Defined in table: Sample
type SampleDefine struct {

	//唯一ID
	ID int64

	//名称
	Name string `自定义tag:"支持go的struct tag"`

	//图标ID
	IconID int32

	//攻击率
	NumericalRate float32

	//物品id
	ItemID int32

	//BuffID
	BuffID []int32

	//类型
	Type ActorType

	//技能ID列表
	SkillID []int32

	//单结构解析
	SingleStruct *Prop

	//字符串结构
	StrStruct []*Prop
}

// Defined in table: Vertical
type PeerData struct {

	// 名字
	Name string

	// 类型
	Type string
}

// Defined in table: Vertical
type VerticalDefine struct {

	//服务器IP
	ServerIP string

	//调试模式
	DebugMode bool

	//客户端人数限制
	ClientLimit int32

	//端
	Peer *PeerData

	Float float32
}

// Defined in table: Exp
type ExpDefine struct {

	//唯一ID
	Level int32

	//经验值
	Exp int32

	//布尔检查
	BoolChecker bool

	//类型
	Type ActorType
}

// Config 访问接口
type ConfigTable struct {

	// 表格原始数据
	Config

	// 索引函数表
	indexFuncByName map[string]func(*ConfigTable)

	// 清空函数表
	clearFuncByName map[string]func(*ConfigTable)

	SampleByID map[int64]*SampleDefine

	SampleByName map[string]*SampleDefine

	ExpByLevel map[int32]*ExpDefine
}

//Vertical
func (self *ConfigTable) GetVertical() *VerticalDefine {
	return self.Vertical[0]
}

// 从json文件加载
func (self *ConfigTable) Load(filename string) error {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	// 生成索引
	for _, v := range self.clearFuncByName {
		v(self)
	}

	err = json.Unmarshal(data, &self.Config)
	if err != nil {
		return err
	}

	// 生成索引
	for _, v := range self.indexFuncByName {
		v(self)
	}

	return nil
}

// 注册外部索引入口, 索引回调, 清空回调
func (self *ConfigTable) RegisterIndexEntry(name string, indexCallback func(*ConfigTable), clearCallback func(*ConfigTable)) {

	if _, ok := self.indexFuncByName[name]; ok {
		panic("duplicate 'Config' table index entry")
	}

	self.indexFuncByName[name] = indexCallback
	self.clearFuncByName[name] = clearCallback
}

// 创建一个Config表读取实例
func NewConfigTable() *ConfigTable {
	return &ConfigTable{

		indexFuncByName: map[string]func(*ConfigTable){

			"Sample": func(tab *ConfigTable) {

				// Sample
				for _, def := range tab.Sample {

					if _, ok := tab.SampleByID[def.ID]; ok {
						panic(fmt.Sprintf("duplicate index in SampleByID: %v", def.ID))
					}

					if _, ok := tab.SampleByName[def.Name]; ok {
						panic(fmt.Sprintf("duplicate index in SampleByName: %v", def.Name))
					}

					tab.SampleByID[def.ID] = def
					tab.SampleByName[def.Name] = def

				}
			},

			"Exp": func(tab *ConfigTable) {

				// Exp
				for _, def := range tab.Exp {

					if _, ok := tab.ExpByLevel[def.Level]; ok {
						panic(fmt.Sprintf("duplicate index in ExpByLevel: %v", def.Level))
					}

					tab.ExpByLevel[def.Level] = def

				}
			},
		},

		clearFuncByName: map[string]func(*ConfigTable){

			"Sample": func(tab *ConfigTable) {

				// Sample

				tab.SampleByID = make(map[int64]*SampleDefine)
				tab.SampleByName = make(map[string]*SampleDefine)
			},

			"Exp": func(tab *ConfigTable) {

				// Exp

				tab.ExpByLevel = make(map[int32]*ExpDefine)
			},
		},

		SampleByID: make(map[int64]*SampleDefine),

		SampleByName: make(map[string]*SampleDefine),

		ExpByLevel: make(map[int32]*ExpDefine),
	}
}