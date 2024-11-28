package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMsdtcLocalHostRequest Request Object
type DeleteMsdtcLocalHostRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeleteMsdtcLocalHostRequestBody `json:"body,omitempty"`
}

func (o DeleteMsdtcLocalHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMsdtcLocalHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteMsdtcLocalHostRequest", string(data)}, " ")
}
