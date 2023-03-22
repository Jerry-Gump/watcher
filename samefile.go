//go:build !windows
// +build !windows

/*
 * @Date: 2023-03-22 12:39:02
 * @LastEditors: Jerry Gump gongzengli@qq.com
 * @LastEditTime: 2023-03-22 12:57:23
 * @FilePath: e:\VSCode-Project\legalsoft.com.cn\watcher\samefile.go
 * @Description:
 * Copyright (c) 2023 by ${git_name} email: ${git_email}, All Rights Reserved.
 */

package watcher

import "os"

func sameFile(fi1, fi2 os.FileInfo) bool {
	return os.SameFile(fi1, fi2)
}
