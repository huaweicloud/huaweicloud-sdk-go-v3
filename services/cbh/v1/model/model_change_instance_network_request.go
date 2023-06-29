package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceNetworkRequest Request Object
type ChangeInstanceNetworkRequest struct {

	// 云堡垒机实例ID，使用UUID格式。
	ServerId string `json:"server_id"`

	Body *ChangeInstanceNetwork `json:"body,omitempty"`
}

func (o ChangeInstanceNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceNetworkRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceNetworkRequest", string(data)}, " ")
}
