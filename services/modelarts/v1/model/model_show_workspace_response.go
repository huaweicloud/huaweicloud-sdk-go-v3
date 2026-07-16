package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceResponse Response Object
type ShowWorkspaceResponse struct {

	// 授权用户列表。默认为空。需要与“auth_type”参数配合使用，且仅当授权类型为“INTERNAL”时才会生效。
	Grants *[]ViewWorkspaceResponseGrants `json:"grants,omitempty"`

	// 创建者名称。
	Owner *string `json:"owner,omitempty"`

	// 授权类型。可选值有PUBLIC、PRIVATE、INTERNAL。默认值为PUBLIC。 - PUBLIC：租户内部公开访问。 - PRIVATE：仅创建者和主账号可访问。 - INTERNAL：创建者、主账号、指定IAM子账号可访问，需要与grants参数配合使用。
	AuthType *string `json:"auth_type,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 最后修改时间，UTC。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创建时间，UTC。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 工作空间名称。长度限制为4-64字符[，支持中文、大小写字母、数字、中划线和下划线](tag:hc,hk)。同时'default'为系统预留的默认工作空间名称，用户无法自己创建名为'default'的工作空间。
	Name *string `json:"name,omitempty"`

	// 工作空间描述，默认为空。长度限制为0-256字符。
	Description *string `json:"description,omitempty"`

	// 工作空间ID，系统生成的32位UUID，不带橫线。
	Id *string `json:"id,omitempty"`

	// 工作空间状态。 - CREATE_FAILED：创建失败。 - NORMAL：状态正常。 - DELETING：正在删除。 - DELETE_FAILED：删除失败。
	Status *string `json:"status,omitempty"`

	// 状态描述，默认为空。该字段会补充显示状态的详细信息。如删除失败时，可通过该字段查看删除失败的原因。
	StatusInfo     *string `json:"status_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceResponse", string(data)}, " ")
}
