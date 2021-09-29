package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateGaussMySqlQuotasResponse struct {
	// 资源列表对象。

	QuotaList      *[]SetQuota `json:"quota_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateGaussMySqlQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlQuotasResponse", string(data)}, " ")
}
