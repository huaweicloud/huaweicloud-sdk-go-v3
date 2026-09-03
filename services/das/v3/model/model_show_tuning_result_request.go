package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTuningResultRequest Request Object
type ShowTuningResultRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ShowTuningResultRequestBody `json:"body,omitempty"`
}

func (o ShowTuningResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTuningResultRequest struct{}"
	}

	return strings.Join([]string{"ShowTuningResultRequest", string(data)}, " ")
}
