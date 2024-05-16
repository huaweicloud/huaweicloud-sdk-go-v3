package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostResultsResponse Response Object
type ListHostResultsResponse struct {

	// 主机漏洞总数
	Total *int32 `json:"total,omitempty"`

	// 主机漏洞信息列表
	Items          *[]HostVulnItem `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListHostResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostResultsResponse struct{}"
	}

	return strings.Join([]string{"ListHostResultsResponse", string(data)}, " ")
}
