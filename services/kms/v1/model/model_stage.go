package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 凭据状态。
type Stage struct {

	// 凭据的版本状态名称。  约束：最小长度1，最大长度64。
	Name *string `json:"name,omitempty" xml:"name"`

	// 凭据的版本状态更新的时间戳，时间戳，即从1970年1月1日至该时间的总秒数。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 凭据的资源标识符。
	SecretId *string `json:"secret_id,omitempty" xml:"secret_id"`

	// 凭据的版本号标识符。
	VersionId *string `json:"version_id,omitempty" xml:"version_id"`
}

func (o Stage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stage struct{}"
	}

	return strings.Join([]string{"Stage", string(data)}, " ")
}
