package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainAllResourceResponse Response Object
type ListDomainAllResourceResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 资源列表
	Resources      *[]CsbResource `json:"resources,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDomainAllResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainAllResourceResponse struct{}"
	}

	return strings.Join([]string{"ListDomainAllResourceResponse", string(data)}, " ")
}
