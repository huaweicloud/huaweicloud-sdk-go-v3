package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EncCertInfo struct {

	// 加密证书序列号。
	EncSerialNumber string `json:"enc_serial_number"`
}

func (o EncCertInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncCertInfo struct{}"
	}

	return strings.Join([]string{"EncCertInfo", string(data)}, " ")
}
