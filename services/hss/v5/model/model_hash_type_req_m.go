package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HashTypeReqM **参数解释**： hash类型 **约束限制**： 必填 **取值范围**： - SHA-256：sha256sum - MD5：md5sum - SHA-1：sha1sum  **默认取值**： 不涉及
type HashTypeReqM struct {
}

func (o HashTypeReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HashTypeReqM struct{}"
	}

	return strings.Join([]string{"HashTypeReqM", string(data)}, " ")
}
