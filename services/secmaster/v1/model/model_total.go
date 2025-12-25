package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Total 总条数
type Total struct {
}

func (o Total) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Total struct{}"
	}

	return strings.Join([]string{"Total", string(data)}, " ")
}
