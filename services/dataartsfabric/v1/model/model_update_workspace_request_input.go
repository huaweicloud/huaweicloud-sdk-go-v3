package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceRequestInput 更新Workspace的请求体。
type UpdateWorkspaceRequestInput struct {

	// 工作空间名称。
	Name *string `json:"name,omitempty"`

	// 描述。用户输入的描述，最长为255个字符。
	Description *string `json:"description,omitempty"`

	// Metastore信息，LakeFormation服务的实例Id，即MetaStoreId。
	MetastoreId *string `json:"metastore_id,omitempty"`
}

func (o UpdateWorkspaceRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceRequestInput struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceRequestInput", string(data)}, " ")
}
