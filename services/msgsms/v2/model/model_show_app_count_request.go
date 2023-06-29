package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppCountRequest Request Object
type ShowAppCountRequest struct {

	// 地域 1. cn：国内 2. intl：国际
	Region string `json:"region"`
}

func (o ShowAppCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppCountRequest struct{}"
	}

	return strings.Join([]string{"ShowAppCountRequest", string(data)}, " ")
}
