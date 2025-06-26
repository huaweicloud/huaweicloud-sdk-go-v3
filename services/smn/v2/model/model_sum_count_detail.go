package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SumCountDetail struct {

	// 短信发送总量
	Sms int32 `json:"sms"`
}

func (o SumCountDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SumCountDetail struct{}"
	}

	return strings.Join([]string{"SumCountDetail", string(data)}, " ")
}
