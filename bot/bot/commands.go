package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func level(s *discordgo.Session, m *discordgo.MessageCreate) {
	emb := discordgo.MessageEmbed{}
	emb.Color = UserList[m.Author.ID].color
	emb.Title = UserList[m.Author.ID].name
	emb.Description = "Class: " + UserList[m.Author.ID].class + "\n" +
		"Level: " + strconv.Itoa(UserList[m.Author.ID].level) + "\n" +
		"XP   : " + strconv.Itoa(UserList[m.Author.ID].xp) + "/" + strconv.Itoa(UserList[m.Author.ID].nxLvl) + "\n"
	embb := &emb

	if _, ok := UserList[m.Author.ID]; ok {
		s.ChannelMessageSendEmbed(m.ChannelID, embb)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please make a Character first with !newCharacter NAME CLASS with NAME being your character's name, and CLASS being your character's class")
	}
}

func newChar(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if len(command) == 3 {
		if _, ok := UserList[m.Author.ID]; !ok {
			UserList[m.Author.ID] = newCharacter(command[1], command[2])
			s.ChannelMessageSend(m.ChannelID, "Congrats, "+command[1]+"! Your Character has been made. Enjoy your journey!")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Sorry! You already have a character, if you would like to remove them use !RemoveCharacter")
		}
	}
}

func ssshhh(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if _, ok := UserList[m.Author.ID]; ok {
		if len(command) == 2 {
			amount, err := strconv.Atoi(command[1])
			if err == nil {
				didlvl := addXP(m.Author.ID, amount)
				s.ChannelMessageSend(m.ChannelID, "Added "+command[1]+" xp to "+UserList[m.Author.ID].name)
				if didlvl == 1 {
					s.ChannelMessageSend(m.ChannelID, UserList[m.Author.ID].name+" Leveled up! Level: "+strconv.Itoa(UserList[m.Author.ID].level))
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, "Please enter a valid integer")
			}
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please Make a character First")
	}
}
func setColor(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if len(command) == 2 {
		if val, ok := UserList[m.Author.ID]; ok {
			num, err := strconv.ParseInt(command[1], 16, 64)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Please send a valid hex code for a color without the leading 0x. Ex: !setcolor 065BBE")
			}
			val.color = int(num)
			UserList[m.Author.ID] = val
		} else {
			s.ChannelMessageSend(m.ChannelID, "Please make a Character first with ```!newCharacter NAME CLASS``` with NAME being your character's name, and CLASS being your character's class")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please send a hex code for a color without the leading 0x. Ex: !setcolor 065BBE")
	}
}

func levels(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := listLvl()
	lvls := discordgo.MessageEmbed{}
	lvls.Color = UserList[m.Author.ID].color
	lvls.Title = UserList[m.Author.ID].name
	lvls.Description = msg
	lvlsemb := &lvls
	s.ChannelMessageSendEmbed(m.ChannelID, lvlsemb)
}

func listLvl() string {
	xp := 19000
	nxLvl := 100
	level := 0
	out := "Level  |         XP\n"
	for xp >= nxLvl {
		level = level + 1
		nxLvl = getNextXP(level)
		out = out + fmt.Sprintf("%-7d|%11d\n", level, nxLvl)
	}
	return out
}

func listClasses(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := discordgo.MessageEmbed{}
	msg.Color = 0
	msg.Title = "Classes!"
	msg.Description = "All classes and their stats"

	for class, struc := range classList {
		stat := ""
		v := reflect.ValueOf(struc)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			stat = stat + fmt.Sprintf("%15s: %5v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}

		feld := discordgo.MessageEmbedField{
			Name:   class,
			Value:  stat,
			Inline: true,
		}
		fld := &feld
		msg.Fields = append(msg.Fields, fld)
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &msg)
}
