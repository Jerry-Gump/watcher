//go:build windows
// +build windows

/*
 * @Date: 2023-03-22 12:39:02
 * @LastEditors: Jerry Gump gongzengli@qq.com
 * @LastEditTime: 2023-03-22 12:58:25
 * @FilePath: e:\VSCode-Project\legalsoft.com.cn\watcher\ishidden_windows.go
 * @Description:
 * Copyright (c) 2023 by ${git_name} email: ${git_email}, All Rights Reserved.
 */

package watcher

import (
	"syscall"
)

func isHiddenFile(path string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}

	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
