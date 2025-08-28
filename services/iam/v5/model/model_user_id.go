package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserId IAM用户ID。
type UserId struct {
}

func (o UserId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserId struct{}"
	}

	return strings.Join([]string{"UserId", string(data)}, " ")
}
