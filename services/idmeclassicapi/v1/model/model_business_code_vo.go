package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BusinessCodeVo struct {

	// 策略集编码。
	Code *string `json:"code,omitempty"`
}

func (o BusinessCodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCodeVo struct{}"
	}

	return strings.Join([]string{"BusinessCodeVo", string(data)}, " ")
}
