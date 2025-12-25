package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBy 创建者
type CreateBy struct {
}

func (o CreateBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBy struct{}"
	}

	return strings.Join([]string{"CreateBy", string(data)}, " ")
}
