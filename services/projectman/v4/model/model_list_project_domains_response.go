package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectDomainsResponse Response Object
type ListProjectDomainsResponse struct {

	// 领域总数
	Total *int32 `json:"total,omitempty"`

	// 领域列表
	Domains        *[]CreateProjectDomainResponseBody `json:"domains,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListProjectDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectDomainsResponse", string(data)}, " ")
}
