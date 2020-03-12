package model

import "fmt"

// GetUserIDTag GetUserIDTag
func GetUserIDTag(id int64) string {
	return fmt.Sprintf("tag_user_id_%d", id)
}

// GetUserEmailTag GetUserEmailTag
func GetUserEmailTag(email string) string {
	return fmt.Sprintf("tag_user_email_%s", email)
}

// GetUserUsernameTag GetUserUsernameTag
func GetUserUsernameTag(username string) string {
	return fmt.Sprintf("tag_user_username_%s", username)
}

// GetUserMobileTag GetUserMobileTag
func GetUserMobileTag(mobile string) string {
	return fmt.Sprintf("tag_user_mobile_%s", mobile)
}

// GetUserAllTag GetUserAllTag
func GetUserAllTag() string {
	return fmt.Sprintf("tag_user_all")
}
