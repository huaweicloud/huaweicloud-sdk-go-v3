package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonTopResponseInfo struct {

	// **参数解释**: 名称 **取值范围**: 字符长度1-128位
	Name *string `json:"name,omitempty"`

	// **参数解释**: 主机数量 **取值范围**: 取值0-2147483647位
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**:  主机占用百分比  **取值范围**:  取值0-100位
	Percentage *int32 `json:"percentage,omitempty"`
}

func (o CommonTopResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonTopResponseInfo struct{}"
	}

	return strings.Join([]string{"CommonTopResponseInfo", string(data)}, " ")
}
