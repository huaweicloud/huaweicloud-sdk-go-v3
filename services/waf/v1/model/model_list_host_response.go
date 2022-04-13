package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostResponse struct {
	// 云模式防护域名的数量

	Total *int32 `json:"total,omitempty"`
	// 详细的防护域名信息

	Items          *[]CloudWafHostItem `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostResponse struct{}"
	}

	return strings.Join([]string{"ListHostResponse", string(data)}, " ")
}
