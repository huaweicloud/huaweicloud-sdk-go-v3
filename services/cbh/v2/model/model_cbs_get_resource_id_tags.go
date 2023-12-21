package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbsGetResourceIdTags TMS查询当前project下所有标签的返回体。
type CbsGetResourceIdTags struct {

	// 执行动作。 - create,创建 - delete,删除
	Action string `json:"action"`

	// 标签列表  租户权限时该字段必选，op_service权限时和sys_tags二选一。
	Tags []ResourceTag `json:"tags"`

	// 系统标签列表  op_service权限可以访问，和tags二选一。  目前TMS调用时只包含一个resource_tag结构体 ，key固定为：_sys_enterprise_project_id  value是UUID或0,value为0表示默认企业项目。  现在仅支持create操作。
	SysTags []ResourceTag `json:"sys_tags"`
}

func (o CbsGetResourceIdTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbsGetResourceIdTags struct{}"
	}

	return strings.Join([]string{"CbsGetResourceIdTags", string(data)}, " ")
}
