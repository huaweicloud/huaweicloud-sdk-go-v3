/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type FileCreate struct {
	// 提交描述。
	Message string `json:"message"`
	// 经base64编码的文件内容。
	Content string `json:"content"`
}

func (o FileCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"FileCreate", string(data)}, " ")
}
