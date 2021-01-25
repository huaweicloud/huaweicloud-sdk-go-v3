package model

import (
	"encoding/json"

	"strings"
)

type StartJobReq struct {
	JobParams *[]JobParam `json:"jobParams,omitempty"`
}

func (o StartJobReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartJobReq struct{}"
	}

	return strings.Join([]string{"StartJobReq", string(data)}, " ")
}
