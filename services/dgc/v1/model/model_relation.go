package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Relation 依赖的job信息
type Relation struct {

	// 依赖类型
	RelationType *string `json:"relationType,omitempty"`

	// 作业名称
	TaskName *string `json:"taskName,omitempty"`

	// 作业ID
	TaskId *int64 `json:"taskId,omitempty"`

	// 脚本名称
	ScriptName *string `json:"scriptName,omitempty"`

	// 脚本ID
	ScriptId *string `json:"scriptId,omitempty"`

	// 用户名
	UserName *string `json:"userName,omitempty"`

	// 版本号
	Version *int32 `json:"version,omitempty"`

	// 依赖类型ID
	RelationId *string `json:"relationId,omitempty"`

	Deleted *bool `json:"deleted,omitempty"`
}

func (o Relation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Relation struct{}"
	}

	return strings.Join([]string{"Relation", string(data)}, " ")
}
