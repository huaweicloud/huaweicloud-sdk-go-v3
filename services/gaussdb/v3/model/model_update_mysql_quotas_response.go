package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMysqlQuotasResponse struct {
	// 资源列表对象。

	QuotaList      *[]SetQuota `json:"quota_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateMysqlQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMysqlQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateMysqlQuotasResponse", string(data)}, " ")
}
