package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapUserId Iam用户ID
type IsapUserId struct {
}

func (o IsapUserId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapUserId struct{}"
	}

	return strings.Join([]string{"IsapUserId", string(data)}, " ")
}
