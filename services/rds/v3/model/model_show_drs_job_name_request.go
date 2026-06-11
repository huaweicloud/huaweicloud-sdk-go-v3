package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrsJobNameRequest Request Object
type ShowDrsJobNameRequest struct {

	// 实例id，源实例
	InstanceId string `json:"instance_id"`
}

func (o ShowDrsJobNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrsJobNameRequest struct{}"
	}

	return strings.Join([]string{"ShowDrsJobNameRequest", string(data)}, " ")
}
