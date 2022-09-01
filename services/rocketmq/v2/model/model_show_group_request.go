package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 消费组名称。
	Group string `json:"group" xml:"group"`
}

func (o ShowGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupRequest", string(data)}, " ")
}
