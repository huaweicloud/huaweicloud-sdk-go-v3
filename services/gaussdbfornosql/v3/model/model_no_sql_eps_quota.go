package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlEpsQuota struct {

	// 实例配额。
	Instance *int32 `json:"instance,omitempty"`

	// vcpus配额。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// ram配额。
	Ram *int32 `json:"ram,omitempty"`
}

func (o NoSqlEpsQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlEpsQuota struct{}"
	}

	return strings.Join([]string{"NoSqlEpsQuota", string(data)}, " ")
}
