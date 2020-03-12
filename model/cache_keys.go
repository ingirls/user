package model

import "fmt"

// GetUserIDKey GetUserIDKey
func GetUserIDKey(id int64) string {
	return fmt.Sprintf("key_user_id:%d", id)
}

// GetUserEmailKey GetUserEmailKey
func GetUserEmailKey(email string) string {
	return fmt.Sprintf("key_user_email:%s", email)
}

// GetUserUsernameKey GetUserUsernameKey
func GetUserUsernameKey(username string) string {
	return fmt.Sprintf("key_user_username:%s", username)
}

// GetUserMobileKey GetUserMobileKey
func GetUserMobileKey(mobile string) string {
	return fmt.Sprintf("key_user_mobile:%s", mobile)
}
