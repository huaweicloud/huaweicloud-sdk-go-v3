package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedDnsRequest Request Object
type ShowRelatedDnsRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	// 恢复时间。
	RestoreTime string `json:"restore_time"`
}

func (o ShowRelatedDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedDnsRequest struct{}"
	}

	return strings.Join([]string{"ShowRelatedDnsRequest", string(data)}, " ")
}
