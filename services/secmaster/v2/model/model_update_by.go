package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBy 更新者
type UpdateBy struct {
}

func (o UpdateBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBy struct{}"
	}

	return strings.Join([]string{"UpdateBy", string(data)}, " ")
}
