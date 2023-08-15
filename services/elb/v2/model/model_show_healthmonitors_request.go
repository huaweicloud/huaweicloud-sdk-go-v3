package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthmonitorsRequest Request Object
type ShowHealthmonitorsRequest struct {

	// 健康检查id
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthmonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsRequest", string(data)}, " ")
}
