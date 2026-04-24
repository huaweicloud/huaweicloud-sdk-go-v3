package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UtcDateTime Timestamp in RFC 3339 format (UTC)
type UtcDateTime struct {
}

func (o UtcDateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UtcDateTime struct{}"
	}

	return strings.Join([]string{"UtcDateTime", string(data)}, " ")
}
