package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceRequestInput 创建Workspace的请求体
type CreateWorkspaceRequestInput struct {

	// 工作空间名称。
	Name string `json:"name"`

	// 描述。用户输入的描述，最长为255个字符。
	Description *string `json:"description,omitempty"`

	// Metastore信息，LakeFormation服务的实例Id，即MetaStoreId。
	MetastoreId *string `json:"metastore_id,omitempty"`

	// 企业项目ID，只有对接了企业项目才可以填写。只能包含字母、数字和中划线，且长度为1到64个字符。默认是0，即default
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateWorkspaceRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceRequestInput struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceRequestInput", string(data)}, " ")
}
