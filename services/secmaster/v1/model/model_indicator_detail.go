package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorDetail 情报详情信息
type IndicatorDetail struct {

	// 威胁情报ID
	Id *string `json:"id,omitempty"`

	// 威胁情报名称
	Name string `json:"name"`

	DataObject *IndicatorDataObjectDetail `json:"data_object,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	DataclassRef *DataClassRefPojo `json:"dataclass_ref,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o IndicatorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDetail struct{}"
	}

	return strings.Join([]string{"IndicatorDetail", string(data)}, " ")
}
