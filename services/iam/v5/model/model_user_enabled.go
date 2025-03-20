package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserEnabled IAM用户是否启用。
type UserEnabled struct {
}

func (o UserEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserEnabled struct{}"
	}

	return strings.Join([]string{"UserEnabled", string(data)}, " ")
}
