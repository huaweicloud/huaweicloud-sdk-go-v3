package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallCbhEipRequest Request Object
type UninstallCbhEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式。
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o UninstallCbhEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallCbhEipRequest struct{}"
	}

	return strings.Join([]string{"UninstallCbhEipRequest", string(data)}, " ")
}
