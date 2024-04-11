package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGetAllParentListRequest Request Object
type ListGetAllParentListRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoQueryParentDto `json:"body,omitempty"`
}

func (o ListGetAllParentListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGetAllParentListRequest struct{}"
	}

	return strings.Join([]string{"ListGetAllParentListRequest", string(data)}, " ")
}
