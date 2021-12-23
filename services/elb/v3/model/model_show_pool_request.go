package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPoolRequest struct {
	// 后端服务器组ID。

	PoolId string `json:"pool_id"`
}

func (o ShowPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}
