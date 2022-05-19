package main

import (
	"fmt"
	"strings"

	tele "gopkg.in/tucnak/telebot.v3"
)

func genCaption(user *tele.User) string {
	desc := "Welcome to the chat 👋" +
		"\n\nBefore you can participate you have to solve our captcha in the next 15 minutes." +
		"\n\nIf you don't complete these steps in time, you will be automatically banned." +
		"\nYou can fail the captcha twice before you are banned." +
		"\n\nJust click the buttons below that match the emojis above!"

	mention := fmt.Sprintf(`[%v](tg://user?id=%v)`, user.FirstName, user.ID)
	caption := fmt.Sprintf("%v, %v", mention, desc)
	return caption
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeRedundantSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func sanitizeName(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	clean := removeRedundantSpaces(result.String())
	return clean
}
