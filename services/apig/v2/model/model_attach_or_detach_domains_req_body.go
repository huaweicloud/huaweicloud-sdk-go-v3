package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 证书批量绑定或解绑域名请求体
type AttachOrDetachDomainsReqBody struct {

	// 证书绑定或解绑域名列表
	Domains []AttachOrDetachDomainInfo `json:"domains"`
}

func (o AttachOrDetachDomainsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachOrDetachDomainsReqBody struct{}"
	}

	return strings.Join([]string{"AttachOrDetachDomainsReqBody", string(data)}, " ")
}
