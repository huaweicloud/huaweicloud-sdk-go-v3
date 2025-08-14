package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupTagResponse Response Object
type ShowServerGroupTagResponse struct {

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`

	// 仅op_service权限才可以获取此字段，非op_service场景不能返回此字段，目前只包含一个resource_tag结构体。 > - key：_sys_enterprise_project_id。 > - value：企业项目id，0表示默认企业项目。
	SysTags        *[]TmsTag `json:"sys_tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowServerGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupTagResponse struct{}"
	}

	return strings.Join([]string{"ShowServerGroupTagResponse", string(data)}, " ")
}
