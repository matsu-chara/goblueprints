package main

import "errors"

// ErrNoAvatarURL はインスタンスがURLを返せないときに発生します
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得出来ません")

// Avatar はユーザーのプロフィール画像を表します
type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar はAvatarの一種です
type AuthAvatar struct{}

// UseAuthAvatar です
var UseAuthAvatar AuthAvatar

// GetAvatarURL です
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar です
type GravatarAvatar struct{}

// UseGravatarAvatar です
var UseGravatarAvatar GravatarAvatar

// GetAvatarURL です
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
