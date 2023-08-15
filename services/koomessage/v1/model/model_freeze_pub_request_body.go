package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezePubRequestBody 服务号冻结，会冻结下属所有资源。不需审核，即时生效。
type FreezePubRequestBody struct {

	// 冻结原因。
	ChangeReason string `json:"change_reason"`
}

func (o FreezePubRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubRequestBody struct{}"
	}

	return strings.Join([]string{"FreezePubRequestBody", string(data)}, " ")
}
