package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTrustedKeySpaceResponseBodyKeyspaces struct {

	// **参数解释：** 密钥空间ID **取值范围：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 密钥空间名称 **取值范围：** 不涉及
	KeyspaceName string `json:"keyspace_name"`

	// **参数解释：** 企业项目ID **取值范围：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 密钥空间类型 **取值范围：** 不涉及
	Type int32 `json:"type"`

	// **参数解释：** 密钥空间状态 **取值范围：** 1：启用，2：冻结
	State int32 `json:"state"`

	// **参数解释：** 密钥空间描述信息 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 密钥空间创建人 **取值范围：** 不涉及
	CreatedBy string `json:"created_by"`

	// **参数解释：** 密钥空间创建时间 **取值范围：** 不涉及
	CreateTime string `json:"create_time"`

	// **参数解释：** 密钥空间最近更新时间 **取值范围：** 不涉及
	LastModifyTime string `json:"last_modify_time"`
}

func (o ListTrustedKeySpaceResponseBodyKeyspaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedKeySpaceResponseBodyKeyspaces struct{}"
	}

	return strings.Join([]string{"ListTrustedKeySpaceResponseBodyKeyspaces", string(data)}, " ")
}
