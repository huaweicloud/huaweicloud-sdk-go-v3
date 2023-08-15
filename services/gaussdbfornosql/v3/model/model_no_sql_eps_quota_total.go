package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlEpsQuotaTotal struct {

	// 实例配额。
	Instance int32 `json:"instance"`

	// vcpus配额。
	Vcpus int32 `json:"vcpus"`

	// ram配额。
	Ram int32 `json:"ram"`
}

func (o NoSqlEpsQuotaTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlEpsQuotaTotal struct{}"
	}

	return strings.Join([]string{"NoSqlEpsQuotaTotal", string(data)}, " ")
}
