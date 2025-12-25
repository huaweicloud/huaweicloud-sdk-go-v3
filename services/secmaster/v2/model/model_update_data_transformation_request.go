package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataTransformationRequest Request Object
type UpdateDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据转换 ID
	DataTransformationId string `json:"data_transformation_id"`

	Body *UpdateDataTransformationRequestBody `json:"body,omitempty"`
}

func (o UpdateDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataTransformationRequest", string(data)}, " ")
}
