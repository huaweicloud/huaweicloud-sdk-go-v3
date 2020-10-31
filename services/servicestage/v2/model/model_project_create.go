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

type ProjectCreate struct {
	// 项目名称。
	Name string `json:"name"`
}

func (o ProjectCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ProjectCreate", string(data)}, " ")
}
