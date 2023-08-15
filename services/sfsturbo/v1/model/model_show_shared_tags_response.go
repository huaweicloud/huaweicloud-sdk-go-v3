package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSharedTagsResponse Response Object
type ShowSharedTagsResponse struct {

	// tag标签的列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 仅op_service权限才可以获取此字段。  1.  目前只包含一个resource_tag结构体 key：_sys_enterprise_project_id  2.  目前key下面只包含一个value，0表示默认企业项目。  非op_service场景不能返回此字段。
	SysTags        *[]ResourceTag `json:"sys_tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSharedTagsResponse", string(data)}, " ")
}
