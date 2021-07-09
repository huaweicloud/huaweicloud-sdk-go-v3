package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlInstanceInfoResponse struct {
	Instance       *MysqlInstanceInfoDetail `json:"instance,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowMysqlInstanceInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlInstanceInfoResponse", string(data)}, " ")
}
