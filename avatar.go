package main

import (
	"errors"
	"io/ioutil"
	"path"
)

//ErrNoAvatar is the error that return ed when the Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

//Avatar reporesents types capable of representing user profile pictures.
type Avatar interface {
	//GetAvatarURL gets the avatar URL for the specified client, or returns an error if sth is wrong.
	GetAvatarURL(c *client) (string, error) //string is the URL if things go well and an error in case sth goes wrong.
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) { //objectなしでもOK
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return  url, nil
}

type FileSystemAvatar struct{}

var UserFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(u chatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(u chatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}
