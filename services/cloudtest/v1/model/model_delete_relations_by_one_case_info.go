package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRelationsByOneCaseInfo struct {
	WorkItemIds *[]string `json:"work_item_ids,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 版本uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 关联关系类型
	RelateType *string `json:"relate_type,omitempty"`
}

func (o DeleteRelationsByOneCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationsByOneCaseInfo struct{}"
	}

	return strings.Join([]string{"DeleteRelationsByOneCaseInfo", string(data)}, " ")
}
