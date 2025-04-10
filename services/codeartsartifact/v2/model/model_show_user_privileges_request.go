package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserPrivilegesRequest Request Object
type ShowUserPrivilegesRequest struct {

	// **参数解释**: 项目id，对应\"需求管理 CodeArts Req\"项目唯一标识，私有依赖库首页地址栏url https://{host}/cloudartifact/project/{project_id}/repository中project_id变量的值。 **约束限制**: 不涉及。 **取值范围**: 只能使用小写英文字符及数字，字符串长度为32位。 **默认取值**: 不涉及。
	ProjectId string `json:"project_id"`
}

func (o ShowUserPrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserPrivilegesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserPrivilegesRequest", string(data)}, " ")
}
