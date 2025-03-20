package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedAt 身份策略的附加时间。
type AttachedAt struct {
}

func (o AttachedAt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedAt struct{}"
	}

	return strings.Join([]string{"AttachedAt", string(data)}, " ")
}
