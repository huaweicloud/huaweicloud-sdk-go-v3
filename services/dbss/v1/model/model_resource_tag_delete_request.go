package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagDeleteRequest struct {

	// 标签列表 租户权限时该字段必选，op_service权限时和sys_tags二选一
	Tags *[]ResourceTagDeleteRequestTags `json:"tags,omitempty"`

	// 系统标签列表 op_service权限可以访问，和tags二选一。 目前TMS调用时只包含一个resource_tag结构体 ，key固定为：_sys_enterprise_project_id value是UUID或0,value为0表示默认企业项目。 现在仅支持create操作。
	SysTags *[]ResourceTagDeleteRequestTags `json:"sys_tags,omitempty"`
}

func (o ResourceTagDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagDeleteRequest struct{}"
	}

	return strings.Join([]string{"ResourceTagDeleteRequest", string(data)}, " ")
}
