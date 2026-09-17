package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdsSlowLogTrendRequest Request Object
type ShowDdsSlowLogTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ShowDdsSlowLogTrendRequestBody `json:"body,omitempty"`
}

func (o ShowDdsSlowLogTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdsSlowLogTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowDdsSlowLogTrendRequest", string(data)}, " ")
}
