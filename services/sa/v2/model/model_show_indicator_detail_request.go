package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndicatorDetailRequest Request Object
type ShowIndicatorDetailRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	// ID of indicator
	IndicatorId string `json:"indicator_id"`
}

func (o ShowIndicatorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndicatorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowIndicatorDetailRequest", string(data)}, " ")
}
