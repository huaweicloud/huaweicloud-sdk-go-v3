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

// 构建工程。
type Build struct {
	Parameters *BuildInfoParameters `json:"parameters,omitempty"`
}

func (o Build) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Build", string(data)}, " ")
}
