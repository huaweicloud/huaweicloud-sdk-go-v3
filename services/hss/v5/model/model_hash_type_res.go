package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HashTypeRes **参数解释**： hash类型 **取值范围**： - SHA-256：sha256sum - MD5：md5sum - SHA-1：sha1sum
type HashTypeRes struct {
}

func (o HashTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HashTypeRes struct{}"
	}

	return strings.Join([]string{"HashTypeRes", string(data)}, " ")
}
