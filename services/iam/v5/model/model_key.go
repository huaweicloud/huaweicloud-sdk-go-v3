package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Key 待删除的标签键。
type Key struct {
}

func (o Key) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Key struct{}"
	}

	return strings.Join([]string{"Key", string(data)}, " ")
}
