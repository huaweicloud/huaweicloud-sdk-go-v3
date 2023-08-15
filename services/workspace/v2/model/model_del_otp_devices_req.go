package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelOtpDevicesReq struct {

	// 待解绑的otp配置id数组。
	OtpIds *[]string `json:"otp_ids,omitempty"`
}

func (o DelOtpDevicesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelOtpDevicesReq struct{}"
	}

	return strings.Join([]string{"DelOtpDevicesReq", string(data)}, " ")
}
