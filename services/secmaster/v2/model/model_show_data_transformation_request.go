package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataTransformationRequest Request Object
type ShowDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据转换 ID
	DataTransformationId string `json:"data_transformation_id"`
}

func (o ShowDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"ShowDataTransformationRequest", string(data)}, " ")
}
