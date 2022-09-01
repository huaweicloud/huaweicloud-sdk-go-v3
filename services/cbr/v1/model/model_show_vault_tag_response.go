package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVaultTagResponse struct {

	// 标签列表 tags中key不重复
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 仅op_service权限才可以获取此字段： 目前只包含一个resource_tag结构体  key：_sys_enterprise_project_id value：企业项目id。0表示默认企业项目 非op_service场景不能返回此字段。
	SysTags        *[]SysTag `json:"sys_tags,omitempty" xml:"sys_tags"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVaultTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultTagResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultTagResponse", string(data)}, " ")
}
