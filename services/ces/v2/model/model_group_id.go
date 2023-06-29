package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupId 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
type GroupId struct {
}

func (o GroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupId struct{}"
	}

	return strings.Join([]string{"GroupId", string(data)}, " ")
}
