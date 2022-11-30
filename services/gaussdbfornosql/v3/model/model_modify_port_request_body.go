package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPortRequestBody struct {

	// 新端口号。端口有效范围为2100~9500，暂不支持8636、8637和8638。
	Port int32 `json:"port"`
}

func (o ModifyPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPortRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyPortRequestBody", string(data)}, " ")
}
