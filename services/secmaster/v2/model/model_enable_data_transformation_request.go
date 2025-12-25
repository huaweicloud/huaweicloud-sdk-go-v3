package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataTransformationRequest Request Object
type EnableDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据转换 ID
	DataTransformationId string `json:"data_transformation_id"`
}

func (o EnableDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"EnableDataTransformationRequest", string(data)}, " ")
}
