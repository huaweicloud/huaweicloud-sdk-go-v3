package model

import (
	"encoding/json"

	"strings"
)

//
type ScopeProjectOption struct {
	// 项目ID，id与name二选一即可。

	Id *string `json:"id,omitempty"`
	// 项目名，id与name二选一即可。

	Name *string `json:"name,omitempty"`

	Domain *ScopeDomainOption `json:"domain,omitempty"`
}

func (o ScopeProjectOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ScopeProjectOption struct{}"
	}

	return strings.Join([]string{"ScopeProjectOption", string(data)}, " ")
}
