package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatorName 创建者名称
type CreatorName struct {
}

func (o CreatorName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatorName struct{}"
	}

	return strings.Join([]string{"CreatorName", string(data)}, " ")
}
