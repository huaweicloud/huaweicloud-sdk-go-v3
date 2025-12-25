package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiRolesUserPermissionsRequest Request Object
type ShowMultiRolesUserPermissionsRequest struct {

	// **参数解释**： 项目ID，对应\"需求管理 CodeArts Req\"项目唯一标识，私有依赖库首页地址栏url https://{host}/cloudartifact/project/{project_id}/repository中project_id变量的值。 **约束限制**： 字符串长度32。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`
}

func (o ShowMultiRolesUserPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiRolesUserPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ShowMultiRolesUserPermissionsRequest", string(data)}, " ")
}
