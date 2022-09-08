package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转发的后端主机组的配置。
type CreateRedirectPoolsConfig struct {

	// 后端主机组的ID。
	PoolId string `json:"pool_id"`

	// 后端主机组的权重。 取值：0-100。
	Weight int32 `json:"weight"`
}

func (o CreateRedirectPoolsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectPoolsConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectPoolsConfig", string(data)}, " ")
}
