package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipCountRespData EIP 数量查询反馈
type EipCountRespData struct {

	// 总体EIP数
	EipTotal *int32 `json:"eip_total,omitempty"`

	// 该账号下所有墙防护EIP总数量
	EipProtected *int32 `json:"eip_protected,omitempty"`

	// 当前防火墙防护EIP数量
	EipProtectedSelf *int32 `json:"eip_protected_self,omitempty"`
}

func (o EipCountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipCountRespData struct{}"
	}

	return strings.Join([]string{"EipCountRespData", string(data)}, " ")
}
