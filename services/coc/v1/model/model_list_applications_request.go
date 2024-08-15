package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationsRequest Request Object
type ListApplicationsRequest struct {

	// id列表
	IdList *[]string `json:"id_list,omitempty"`

	// parent id
	ParentId *string `json:"parent_id,omitempty"`

	// 应用名称模糊匹配
	NameLike *string `json:"name_like,omitempty"`

	// 应用code
	Code *string `json:"code,omitempty"`

	// 分页参数，上一页请求最后一个id
	Marker *string `json:"marker,omitempty"`

	// 最大的返回数量
	Limit int32 `json:"limit"`
}

func (o ListApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationsRequest", string(data)}, " ")
}
