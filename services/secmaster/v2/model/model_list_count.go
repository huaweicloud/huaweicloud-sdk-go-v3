package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCount 数量
type ListCount struct {
}

func (o ListCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCount struct{}"
	}

	return strings.Join([]string{"ListCount", string(data)}, " ")
}
