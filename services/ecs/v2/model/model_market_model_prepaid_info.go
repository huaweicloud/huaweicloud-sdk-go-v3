package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MarketModelPrepaidInfo struct {
	ExpiredTime *string `json:"expired_time,omitempty"`
}

func (o MarketModelPrepaidInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarketModelPrepaidInfo struct{}"
	}

	return strings.Join([]string{"MarketModelPrepaidInfo", string(data)}, " ")
}
