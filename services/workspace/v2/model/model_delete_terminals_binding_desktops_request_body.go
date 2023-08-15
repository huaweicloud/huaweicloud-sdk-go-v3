package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTerminalsBindingDesktopsRequestBody struct {

	// 绑定策略ID列表
	IdList *[]string `json:"id_list,omitempty"`
}

func (o DeleteTerminalsBindingDesktopsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTerminalsBindingDesktopsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTerminalsBindingDesktopsRequestBody", string(data)}, " ")
}
