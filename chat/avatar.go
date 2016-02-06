package main

import (
	"errors"
	"path/filepath"
)

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

// FileSystemAvatar .
type FileSystemAvatar struct{}

// UseFileSystemAvatar .
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL .
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	userid, ok := c.userData["userid"]
	if !ok {
		return "", ErrNoAvatarURL
	}
	useridStr, ok := userid.(string)
	if !ok {
		return "", ErrNoAvatarURL
	}

	matches, err := filepath.Glob(filepath.Join("avatars", useridStr+"*"))
	if err != nil || len(matches) == 0 {
		return "", ErrNoAvatarURL
	}
	return "/" + matches[0], nil
}
