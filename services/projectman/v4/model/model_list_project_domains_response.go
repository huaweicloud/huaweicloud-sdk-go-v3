package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectDomainsResponse struct {

	// 领域总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 领域列表
	Domains        *[]CreateProjectDomainResponseBody `json:"domains,omitempty" xml:"domains"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListProjectDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectDomainsResponse", string(data)}, " ")
}
