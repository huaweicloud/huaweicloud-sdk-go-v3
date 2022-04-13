package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TagResource struct {
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 资源详情

	ResourceDetail []Vault `json:"resource_detail"`
	// 标签列表 没有标签默认为空数字。

	Tags []Tag `json:"tags"`
	// 资源名称

	ResourceName string `json:"resource_name"`
	// 仅op_service权限才可以可以获取此字段：  目前只包含一个resource_tag 结构体。  key：_sys_enterprise_project_id  value：企业项目id，0表示默认企业项目  非op_service场景不能返回此字段。

	SysTags []SysTag `json:"sys_tags"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
