package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceGroupId 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
type ResourceGroupId struct {
}

func (o ResourceGroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupId struct{}"
	}

	return strings.Join([]string{"ResourceGroupId", string(data)}, " ")
}
