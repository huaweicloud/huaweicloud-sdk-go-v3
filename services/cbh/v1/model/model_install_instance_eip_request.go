package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallInstanceEipRequest Request Object
type InstallInstanceEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o InstallInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"InstallInstanceEipRequest", string(data)}, " ")
}
