package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SetMysqlQuotasResponse struct {
	// 资源列表对象。

	QuotaList      *[]SetQuota `json:"quota_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o SetMysqlQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetMysqlQuotasResponse struct{}"
	}

	return strings.Join([]string{"SetMysqlQuotasResponse", string(data)}, " ")
}
