package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsResponse Response Object
type ListEnterpriseProjectsResponse struct {

	// 企业项目信息
	Data *[]EpsInfo `json:"data,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnterpriseProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsResponse", string(data)}, " ")
}
