package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociatedResourceRulesRequest Request Object
type ListAssociatedResourceRulesRequest struct {

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页位置标识（索引）。从marker指定索引的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 规则的配置名称
	SettingName *string `json:"setting_name,omitempty"`

	// 规则的区域ID
	RegionId *string `json:"region_id,omitempty"`
}

func (o ListAssociatedResourceRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedResourceRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAssociatedResourceRulesRequest", string(data)}, " ")
}
