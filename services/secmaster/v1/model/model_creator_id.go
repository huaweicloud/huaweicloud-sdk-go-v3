package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatorId 创建者id
type CreatorId struct {
}

func (o CreatorId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatorId struct{}"
	}

	return strings.Join([]string{"CreatorId", string(data)}, " ")
}
