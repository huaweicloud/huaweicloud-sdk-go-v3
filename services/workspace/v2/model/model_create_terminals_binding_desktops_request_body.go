package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTerminalsBindingDesktopsRequestBody struct {

	// 需要新增的MAC绑定VM策略信息列表
	BindList *[]CreateTerminalsBindingDesktopsInfo `json:"bind_list,omitempty"`
}

func (o CreateTerminalsBindingDesktopsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTerminalsBindingDesktopsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTerminalsBindingDesktopsRequestBody", string(data)}, " ")
}
