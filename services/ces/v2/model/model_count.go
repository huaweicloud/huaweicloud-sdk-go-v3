package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Count 次数
type Count struct {
}

func (o Count) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Count struct{}"
	}

	return strings.Join([]string{"Count", string(data)}, " ")
}
