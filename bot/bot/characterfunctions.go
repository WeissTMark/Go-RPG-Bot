package main

// Class Functions:

func addXP(id string, amount int) int {
	user := UserList[id]
	user.xp += amount
	ret := 0
	for user.xp >= user.nxLvl {
		user.level = user.level + 1
		user.nxLvl = getNextXP(user.level)
		ret = 1
	}
	UserList[id] = user
	return ret
}

func newCharacter(name string, class string) Character {
	var c = Character{}
	c.name = name
	c.class = class
	c.level = 0
	c.xp = 0
	c.nxLvl = 100
	c.color = 416702
	return c
}

func setClassDefaults() map[string]Class {
	classes := map[string]Class{}
	classes["Monk"] = Class{79, 67, 79, 65, 53, 65, 55, 1, 1, 11, 36, 30}
	classes["Blacksmith"] = Class{81, 85, 45, 57, 77, 61, 55, 1, 0, 9, 38, 30}
	classes["Hunter"] = Class{57, 71, 51, 83, 69, 83, 55, 1, 1, 20, 37, 30}
	classes["Minstrel"] = Class{49, 55, 75, 73, 83, 71, 55, 0, 1, 13, 35, 30}
	classes["Scholar"] = Class{47, 65, 83, 71, 75, 75, 55, 0, 1, 15, 36, 30}
	classes["Busker"] = Class{77, 73, 53, 59, 81, 63, 55, 1, 1, 10, 37, 30}
	classes["Herbalist"] = Class{49, 57, 81, 75, 63, 69, 55, 0, 1, 13, 35, 30}
	classes["Trapper"] = Class{53, 65, 47, 81, 77, 79, 55, 1, 0, 17, 36, 30}
	classes["Woodcutter"] = Class{83, 79, 49, 75, 57, 63, 55, 1, 0, 10, 37, 30}
	classes["Hobo"] = Class{75, 75, 75, 75, 75, 75, 55, 1, 1, 15, 37, 30}
	classes["Monk"] = Class{79, 67, 79, 65, 53, 65, 55, 1, 1, 11, 36, 30}
	classes["Treasure Hunter"] = Class{47, 61, 79, 79, 79, 63, 55, 0, 1, 10, 36, 30}
	classes["Astronomer"] = Class{55, 79, 81, 73, 51, 67, 55, 1, 1, 12, 37, 30}
	classes["Gladiator"] = Class{81, 57, 45, 81, 55, 75, 55, 1, 0, 15, 35, 30}
	return classes
}

func getNextXP(currLvl int) int {
	return 1
}
