package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndexUsageSwitchNewRequest Request Object
type ShowIndexUsageSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowIndexUsageSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndexUsageSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowIndexUsageSwitchNewRequest", string(data)}, " ")
}
