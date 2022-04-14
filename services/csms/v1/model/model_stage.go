package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 凭据状态。
type Stage struct {
	// 凭据的版本状态名称。  约束：最小长度1，最大长度64。

	Name *string `json:"name,omitempty"`
	// 凭据的版本状态更新的时间戳，时间戳，即从1970年1月1日至该时间的总秒数。

	UpdateTime *int64 `json:"update_time,omitempty"`
	// 凭据名称。

	SecretName *string `json:"secret_name,omitempty"`
	// 凭据的版本号标识符。

	VersionId *string `json:"version_id,omitempty"`
}

func (o Stage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stage struct{}"
	}

	return strings.Join([]string{"Stage", string(data)}, " ")
}
