package main

import (
	"errors"
	"path/filepath"
)

// ErrNoAvatarURL はインスタンスがURLを返せないときに発生します
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得出来ません")

// Avatar はユーザーのプロフィール画像を表します
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}
type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar はAvatarの一種です
type AuthAvatar struct{}

// UseAuthAvatar です
var UseAuthAvatar AuthAvatar

// GetAvatarURL です
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar です
type GravatarAvatar struct{}

// UseGravatarAvatar です
var UseGravatarAvatar GravatarAvatar

// GetAvatarURL です
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar .
type FileSystemAvatar struct{}

// UseFileSystemAvatar .
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL .
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	matches, err := filepath.Glob(filepath.Join("avatars", u.UniqueID()+"*"))
	if err != nil || len(matches) == 0 {
		return "", ErrNoAvatarURL
	}
	return "/" + matches[0], nil
}
