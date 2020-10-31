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

type ComponentModify struct {
	// 应用组件名称。
	Name *string `json:"name,omitempty"`
	// 描述。
	Description *string       `json:"description,omitempty"`
	Source      *SourceObject `json:"source,omitempty"`
	Build       *Build        `json:"build,omitempty"`
}

func (o ComponentModify) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ComponentModify", string(data)}, " ")
}
