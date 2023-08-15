package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultSet struct {

	// 可用性
	Availability *int32 `json:"availability,omitempty"`

	// 类别
	Category *string `json:"category,omitempty"`

	// 作业编号
	JobId *string `json:"job_id,omitempty"`

	// 作业类型
	JobType *string `json:"job_type,omitempty"`

	// rec编号
	RecId *string `json:"rec_id,omitempty"`

	// rec类型
	RecType *string `json:"rec_type,omitempty"`

	// 场景编号
	SceneId *string `json:"scene_id,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 工作空间编号
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ResultSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultSet struct{}"
	}

	return strings.Join([]string{"ResultSet", string(data)}, " ")
}
