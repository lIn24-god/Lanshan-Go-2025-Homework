package main

import (
	"fmt"
)

//==============================基础乐队信息与方法================================

type BaseBand struct {
	characters []Character
	name       string
}

func (ba *BaseBand) AddCharacter(character Character) {
	ba.characters = append(ba.characters, character)
	fmt.Printf("%s 已就位!!!!!\n", character.name)
}
func (ba *BaseBand) GetInformation(targetName string) (Character, bool) {
	for _, character := range ba.characters {
		if character.name == targetName {
			return character, true
		}
	}
	return Character{}, false
}
func (ba *BaseBand) UpdateInterest(targetName, message string) {
	for i := range ba.characters {
		if ba.characters[i].name == targetName {
			current := ba.characters[i]
			ba.characters[i].interest = message
			fmt.Printf("已从 %s 转为 %s\n", current.interest, message)
			fmt.Println("信息已修改!!!!!")
		}
	}
}
func (ba *BaseBand) ShowAll() {
	fmt.Println("===============================乐队信息=================================")
	for _, character := range ba.characters {

		fmt.Printf("%8s  %8s  %8d  %8s  %8s\n", character.name, character.sex, character.age, character.role, character.interest)
	}
	fmt.Println("===================================================================")
}

//============================================================================

type Bander interface {
	AddCharacter(character Character)
	GetInformation(targetName string) (Character, bool)
	UpdateInterest(targetName, message string)
	ShowAll()
}
type Character struct {
	name     string
	sex      string
	role     string
	age      int
	interest string
}

// ================================================具体乐队信息========================================
//组合实现

type Mygo struct {
	BaseBand
	feature string
}
type Mujica struct {
	BaseBand
	style string
}

//============================通用乐队管理函数=============================

func AddMemberToBand(b Bander, name, sex, role string, age int, interest string) {
	character := Character{name, sex, role, age, interest}
	b.AddCharacter(character)
}
func QueryInfo(b Bander, name string) {
	character, found := b.GetInformation(name)
	if !found {
		fmt.Println("未找到成员!!!!!")
		return
	}
	fmt.Printf("===================%s的详细信息=======================\n", character.name)
	fmt.Printf("性别:%s\n", character.sex)
	fmt.Printf("角色:%s\n", character.role)
	fmt.Printf("年龄:%d\n", character.age)
	fmt.Printf("兴趣:%s\n", character.interest)
	fmt.Println("====================================================")
}

//===================================泛型与类型约束===================================
//过滤器

func FilterBy[T any](items []T, shouldInclude func(T) bool) []T {
	var result []T
	for _, item := range items {
		if shouldInclude(item) {
			result = append(result, item)
		}
	}
	return result
}

// ============================主函数================================
func main() {
	Itsmygo := &Mygo{
		BaseBand: BaseBand{
			characters: nil,
			name:       "ItsMyGO!!!!!",
		},
		feature: "底边乐队",
	}
	Avemujica := &Mujica{
		BaseBand: BaseBand{
			characters: nil,
			name:       "AveMujica",
		},
		style: "银河战舰",
	}
	Itsmygo.AddCharacter(Character{"千早爱音", "奥特曼", "吉他手", 15, "soyorin~~~"})
	Itsmygo.AddCharacter(Character{"长崎素世", "16岁离异寡妇JK", "贝斯手", 16, "为什么要演奏春日影!?"})
	Itsmygo.AddCharacter(Character{"椎名立希", "女", "鼓手", 15, "竟敢无视灯"})
	Itsmygo.AddCharacter(Character{"要乐奈", "哈基米", "吉他手", 14, "抹茶芭菲 suki"})
	AddMemberToBand(Itsmygo, "高松灯", "企鹅", "主唱", 15, "咕咕嘎嘎")
	AddMemberToBand(Avemujica, "丰川祥子", "卡密", "键盘手", 15, "你这家伙...真是满脑子只想着自己呢")
	AddMemberToBand(Avemujica, "若叶睦", "睦子米", "吉他手", 15, "祥,移动")
	AddMemberToBand(Avemujica, "三角初华", "女(tong)", "主唱", 15, "saki酱●▛▙saki酱●▛▙saki酱●▛▙")
	AddMemberToBand(Avemujica, "祐天寺若麦", "哈气哈基米", "鼓手", 15, "扣尼及喵姆喵姆")
	AddMemberToBand(Avemujica, "八幡海铃", "职业雇佣兵", "贝斯手", 16, "信任这一块")
	QueryInfo(Itsmygo, "丰川祥子")
	QueryInfo(Avemujica, "丰川祥子")
	QueryInfo(Itsmygo, "高松灯")
	Avemujica.UpdateInterest("丰川祥子", "好想成为人类啊!!!!!")
	Avemujica.ShowAll()

	//过滤器使用

	guitarist := FilterBy(Itsmygo.characters, func(c Character) bool {
		return c.role == "吉他手"
	})
	fmt.Println(guitarist)

	above15 := FilterBy(Itsmygo.characters, func(c Character) bool {
		return c.age >= 15
	})
	fmt.Println(above15)

}
