package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociatedResourceSettingsResponse Response Object
type ListAssociatedResourceSettingsResponse struct {

	// 规则的配置信息
	Settings *[]AssociatedResourceSetting `json:"settings,omitempty"`

	// 记录总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAssociatedResourceSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedResourceSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListAssociatedResourceSettingsResponse", string(data)}, " ")
}
