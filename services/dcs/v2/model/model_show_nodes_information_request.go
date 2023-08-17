package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesInformationRequest Request Object
type ShowNodesInformationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowNodesInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesInformationRequest struct{}"
	}

	return strings.Join([]string{"ShowNodesInformationRequest", string(data)}, " ")
}
