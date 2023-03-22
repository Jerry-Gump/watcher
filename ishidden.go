//go:build !windows
// +build !windows

/*
 * @Date: 2023-03-22 12:39:02
 * @LastEditors: Jerry Gump gongzengli@qq.com
 * @LastEditTime: 2023-03-22 12:56:50
 * @FilePath: e:\VSCode-Project\legalsoft.com.cn\watcher\ishidden.go
 * @Description:
 * Copyright (c) 2023 by ${git_name} email: ${git_email}, All Rights Reserved.
 */

package watcher

import (
	"path/filepath"
	"strings"
)

func isHiddenFile(path string) (bool, error) {
	return strings.HasPrefix(filepath.Base(path), "."), nil
}
