package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息
type QueryQuotaInfo struct {
	Resource *QuotaResource `json:"resource,omitempty" xml:"resource"`
}

func (o QueryQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryQuotaInfo struct{}"
	}

	return strings.Join([]string{"QueryQuotaInfo", string(data)}, " ")
}
