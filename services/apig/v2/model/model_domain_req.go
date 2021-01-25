package model

import (
	"encoding/json"

	"strings"
)

type DomainReq struct {
	// 自定义域名。长度为0-255位的字符串，需要符合域名规范。
	UrlDomain string `json:"url_domain"`
}

func (o DomainReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainReq struct{}"
	}

	return strings.Join([]string{"DomainReq", string(data)}, " ")
}
