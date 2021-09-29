package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlInstanceInfoResponse struct {
	Instance       *MysqlInstanceInfoDetail `json:"instance,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowGaussMySqlInstanceInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlInstanceInfoResponse", string(data)}, " ")
}
