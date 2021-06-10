package model

import (
	"encoding/json"

	"strings"
)

// 域名对象
type CreateDomainRequestBody struct {
	Domain *DomainBody `json:"domain"`
}

func (o CreateDomainRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDomainRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDomainRequestBody", string(data)}, " ")
}
