package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScriptTaskInfoTaskScripts struct {

	// 脚本ID。
	Id *string `json:"id,omitempty"`

	// 脚本名称。
	Name *string `json:"name,omitempty"`
}

func (o ScriptTaskInfoTaskScripts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptTaskInfoTaskScripts struct{}"
	}

	return strings.Join([]string{"ScriptTaskInfoTaskScripts", string(data)}, " ")
}
