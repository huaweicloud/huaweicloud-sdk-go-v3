package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowBlockRequest Request Object
type ShowFlowBlockRequest struct {

	// instanceId
	InstanceId string `json:"instance_id"`
}

func (o ShowFlowBlockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowBlockRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowBlockRequest", string(data)}, " ")
}
