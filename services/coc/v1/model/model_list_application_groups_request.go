package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationGroupsRequest Request Object
type ListApplicationGroupsRequest struct {

	// 分组id集合。
	IdList *[]string `json:"id_list,omitempty"`

	// 组件id。
	ComponentId *string `json:"component_id,omitempty"`

	// 应用id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 分组名称模糊匹配。
	NameLike *string `json:"name_like,omitempty"`

	// 分组code。
	Code *string `json:"code,omitempty"`

	// 分页参数，上一页请求最后一个id。
	Marker *string `json:"marker,omitempty"`

	// 最大的返回数量。
	Limit int32 `json:"limit"`
}

func (o ListApplicationGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationGroupsRequest", string(data)}, " ")
}
