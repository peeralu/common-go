package util

import "strings"

func MarkMobilePhone(phone string) string {
	if len(phone) < 10 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

func MarkEmail(email string) string {
	if len(email) < 5 {
		return email
	}
	atIndex := strings.Index(email, "@")
	if atIndex < 0 {
		return email
	}
	if atIndex < 3 {
		return email[:atIndex] + "****" + email[atIndex:]
	}
	return email[:3] + "****" + email[atIndex:]
}

func MarkIDCard(id string) string {
	if len(id) < 13 {
		return id
	}
	return id[:3] + "****" + id[7:]
}

func MarkBankAccount(account string) string {
	if len(account) < 10 {
		return account
	}
	return account[:3] + "****" + account[7:]
}
