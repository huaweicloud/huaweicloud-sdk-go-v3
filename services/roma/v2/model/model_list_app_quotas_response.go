package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppQuotasResponse struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 客户端配额列表

	Quotas         *[]AppQuotaInfo `json:"quotas,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAppQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListAppQuotasResponse", string(data)}, " ")
}
