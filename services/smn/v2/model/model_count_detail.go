package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountDetail struct {

	// 短信发送数量
	Sms int32 `json:"sms"`
}

func (o CountDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDetail struct{}"
	}

	return strings.Join([]string{"CountDetail", string(data)}, " ")
}
