package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlQuotasResponse struct {
	// 资源列表对象。

	QuotaList *[]Quota `json:"quota_list,omitempty"`
	// 配额记录的条数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMysqlQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlQuotasResponse", string(data)}, " ")
}
