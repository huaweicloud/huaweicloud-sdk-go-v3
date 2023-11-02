package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndicatorsRequest Request Object
type ListIndicatorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *IndicatorListSearchRequest `json:"body,omitempty"`
}

func (o ListIndicatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndicatorsRequest struct{}"
	}

	return strings.Join([]string{"ListIndicatorsRequest", string(data)}, " ")
}
