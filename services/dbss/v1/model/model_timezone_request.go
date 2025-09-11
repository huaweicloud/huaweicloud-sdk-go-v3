package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimezoneRequest struct {

	// 时区信息
	Timezone string `json:"timezone"`
}

func (o TimezoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimezoneRequest struct{}"
	}

	return strings.Join([]string{"TimezoneRequest", string(data)}, " ")
}
