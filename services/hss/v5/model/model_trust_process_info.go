package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrustProcessInfo struct {

	// **参数解释**: 进程路径 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释**: 进程hash **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Hash *string `json:"hash,omitempty"`
}

func (o TrustProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustProcessInfo struct{}"
	}

	return strings.Join([]string{"TrustProcessInfo", string(data)}, " ")
}
