package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePostgresqlInstanceAliasResponse struct {
	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePostgresqlInstanceAliasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlInstanceAliasResponse struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlInstanceAliasResponse", string(data)}, " ")
}
