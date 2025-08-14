package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecuritySensitiveDataDetailsResponse Response Object
type ListSecuritySensitiveDataDetailsResponse struct {

	// 敏感数据发现详情总条数。
	Total *int64 `json:"total,omitempty"`

	// 敏感数据发现列表。
	Content *[]SensitiveDataDto `json:"content,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecuritySensitiveDataDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySensitiveDataDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSecuritySensitiveDataDetailsResponse", string(data)}, " ")
}
