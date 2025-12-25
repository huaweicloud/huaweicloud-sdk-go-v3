package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataTransformationMetricsRequest Request Object
type ListDataTransformationMetricsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListDataTransformationMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataTransformationMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListDataTransformationMetricsRequest", string(data)}, " ")
}
