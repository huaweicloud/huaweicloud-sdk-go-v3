package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyVersionId 身份策略版本号，以\"v\"开头后跟数字，例如\"v5\"。
type PolicyVersionId struct {
}

func (o PolicyVersionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyVersionId struct{}"
	}

	return strings.Join([]string{"PolicyVersionId", string(data)}, " ")
}
