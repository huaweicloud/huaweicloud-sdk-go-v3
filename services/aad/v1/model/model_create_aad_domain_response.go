package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAadDomainResponse Response Object
type CreateAadDomainResponse struct {

	// 高防提供的CNAME地址
	Cname *string `json:"cname,omitempty"`

	// 域名id
	DomainId       *string `json:"domainId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAadDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAadDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateAadDomainResponse", string(data)}, " ")
}
