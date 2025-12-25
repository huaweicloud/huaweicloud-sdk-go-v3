package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDataTransformationRequest Request Object
type DisableDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据转换 ID
	DataTransformationId string `json:"data_transformation_id"`
}

func (o DisableDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"DisableDataTransformationRequest", string(data)}, " ")
}
