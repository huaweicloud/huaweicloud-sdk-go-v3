package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndicatorDetailRequest Request Object
type ShowIndicatorDetailRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 威胁情报ID
	IndicatorId string `json:"indicator_id"`
}

func (o ShowIndicatorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndicatorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowIndicatorDetailRequest", string(data)}, " ")
}
