package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociatedResourceSettingsRequest Request Object
type ListAssociatedResourceSettingsRequest struct {

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页位置标识（索引）。从marker指定索引的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 规则的区域ID
	RegionId *string `json:"region_id,omitempty"`
}

func (o ListAssociatedResourceSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedResourceSettingsRequest struct{}"
	}

	return strings.Join([]string{"ListAssociatedResourceSettingsRequest", string(data)}, " ")
}
