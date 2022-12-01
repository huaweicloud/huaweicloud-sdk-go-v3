package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EIP 数量查询反馈
type EipCountRespData struct {

	// 防护对象ID
	ObjectId *string `json:"object_id,omitempty"`

	// EIP总数
	EipTotal *int32 `json:"eip_total,omitempty"`

	// EIP防护数
	EipProtected *int32 `json:"eip_protected,omitempty"`
}

func (o EipCountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipCountRespData struct{}"
	}

	return strings.Join([]string{"EipCountRespData", string(data)}, " ")
}
