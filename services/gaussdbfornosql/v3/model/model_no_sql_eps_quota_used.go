package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlEpsQuotaUsed struct {

	// 已使用实例配额。
	Instance int32 `json:"instance"`

	// 已使用vcpus配额。
	Vcpus int32 `json:"vcpus"`

	// 已使用ram配额。
	Ram int32 `json:"ram"`
}

func (o NoSqlEpsQuotaUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlEpsQuotaUsed struct{}"
	}

	return strings.Join([]string{"NoSqlEpsQuotaUsed", string(data)}, " ")
}
