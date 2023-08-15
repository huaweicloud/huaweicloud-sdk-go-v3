package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoSqlEpsQuotaRequestInfo struct {

	// 实例配额。
	Instance *int32 `json:"instance,omitempty"`

	// vcpus配额。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// ram配额。
	Ram *int32 `json:"ram,omitempty"`
}

func (o NoSqlEpsQuotaRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlEpsQuotaRequestInfo struct{}"
	}

	return strings.Join([]string{"NoSqlEpsQuotaRequestInfo", string(data)}, " ")
}
