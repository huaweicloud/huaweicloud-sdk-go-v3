package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultSet struct {

	// 可用性
	Availability *int32 `json:"availability,omitempty" xml:"availability"`

	// 类别
	Category *string `json:"category,omitempty" xml:"category"`

	// 作业编号
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 作业类型
	JobType *string `json:"job_type,omitempty" xml:"job_type"`

	// rec编号
	RecId *string `json:"rec_id,omitempty" xml:"rec_id"`

	// rec类型
	RecType *string `json:"rec_type,omitempty" xml:"rec_type"`

	// 场景编号
	SceneId *string `json:"scene_id,omitempty" xml:"scene_id"`

	// 表名
	TableName *string `json:"table_name,omitempty" xml:"table_name"`

	// 工作空间编号
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id"`
}

func (o ResultSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultSet struct{}"
	}

	return strings.Join([]string{"ResultSet", string(data)}, " ")
}
