package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallCbhEipRequest Request Object
type InstallCbhEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o InstallCbhEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallCbhEipRequest struct{}"
	}

	return strings.Join([]string{"InstallCbhEipRequest", string(data)}, " ")
}
