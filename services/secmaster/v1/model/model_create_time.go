package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTime 创建时间
type CreateTime struct {
}

func (o CreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTime struct{}"
	}

	return strings.Join([]string{"CreateTime", string(data)}, " ")
}
