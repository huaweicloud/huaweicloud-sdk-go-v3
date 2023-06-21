package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务号解冻请求体。
type UnfreezePubRequestBody struct {

	// 解冻原因。
	ChangeReason string `json:"change_reason"`
}

func (o UnfreezePubRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezePubRequestBody struct{}"
	}

	return strings.Join([]string{"UnfreezePubRequestBody", string(data)}, " ")
}
