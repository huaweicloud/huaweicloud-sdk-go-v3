package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MfaRsp struct {

	// mfa方法
	Method *string `json:"method,omitempty"`

	// mfa信息
	Info *string `json:"info,omitempty"`
}

func (o MfaRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaRsp struct{}"
	}

	return strings.Join([]string{"MfaRsp", string(data)}, " ")
}
