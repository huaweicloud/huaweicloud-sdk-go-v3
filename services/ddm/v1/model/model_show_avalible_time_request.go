package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleTimeRequest Request Object
type ShowAvalibleTimeRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowAvalibleTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowAvalibleTimeRequest", string(data)}, " ")
}
