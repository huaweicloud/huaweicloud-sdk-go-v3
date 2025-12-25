package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataTransformationRequest Request Object
type DeleteDataTransformationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据转换 ID
	DataTransformationId string `json:"data_transformation_id"`
}

func (o DeleteDataTransformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataTransformationRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataTransformationRequest", string(data)}, " ")
}
