package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlertGroupsByConditionRequestBody struct {

	// 告警组ID
	GroupId *string `json:"group_id,omitempty"`

	// 告警组ID列表
	GroupIds *[]string `json:"group_ids,omitempty"`

	// 当前页数
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 服务ID
	TestServiceId *string `json:"test_service_id,omitempty"`

	// 用户ID列表
	UserIds *[]string `json:"userIds,omitempty"`

	// 用户名
	UserName *string `json:"userName,omitempty"`
}

func (o ListAlertGroupsByConditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertGroupsByConditionRequestBody struct{}"
	}

	return strings.Join([]string{"ListAlertGroupsByConditionRequestBody", string(data)}, " ")
}
