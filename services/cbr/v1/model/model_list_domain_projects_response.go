package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainProjectsResponse Response Object
type ListDomainProjectsResponse struct {

	// 项目列表
	Projects       *[]DomainProjectsInfo `json:"projects,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDomainProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainProjectsResponse", string(data)}, " ")
}
