package main

import (
	"errors"
)

//ErrNoAvatar is the error that return ed when the Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")
//Avatar reporesents types capable of representing user profile pictures.
type Avatar interface {
	//GetAvatarURL gets the avatar URL for the specified client, or returns an error if sth is wrong.
	GetAvatarURL(c *client) (string, error) //string is the URL if things go well and an error in case sth goes wrong.
}

type AuthAvatar struct {}
var UserAuthAvatar AuthAvatar
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok{
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}




