package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrorNoAvatarURL when no value present.")
	}
	//set a value
	testUrl := "http://url-to-getavatar"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present.")
	}
	if url != testUrl {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL.")
	}
}