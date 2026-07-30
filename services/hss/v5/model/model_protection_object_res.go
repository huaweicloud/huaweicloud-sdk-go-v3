package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectionObjectRes **参数解释**： 防护对象 **取值范围**： 字符长度1-128位
type ProtectionObjectRes struct {
}

func (o ProtectionObjectRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionObjectRes struct{}"
	}

	return strings.Join([]string{"ProtectionObjectRes", string(data)}, " ")
}
