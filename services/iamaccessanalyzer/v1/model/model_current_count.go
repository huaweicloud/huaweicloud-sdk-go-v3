package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CurrentCount 当前页中的项数。
type CurrentCount struct {
}

func (o CurrentCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CurrentCount struct{}"
	}

	return strings.Join([]string{"CurrentCount", string(data)}, " ")
}
