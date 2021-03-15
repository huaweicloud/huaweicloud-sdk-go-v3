package model

import (
	"encoding/json"

	"strings"
)

// project信息
type ProjectInfo struct {
	Domain *DomainInfo `json:"domain,omitempty"`
	// project id
	Id *string `json:"id,omitempty"`
	// project name
	Name string `json:"name"`
}

func (o ProjectInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProjectInfo struct{}"
	}

	return strings.Join([]string{"ProjectInfo", string(data)}, " ")
}
