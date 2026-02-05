package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDpeClassifyRequestBody 创建分类映射请求体
type CreateDpeClassifyRequestBody struct {

	// 映射id
	Id string `json:"id"`

	// 名称
	Name string `json:"name"`

	ProjectId *interface{} `json:"project_id,omitempty"`

	WorkspaceId *interface{} `json:"workspace_id,omitempty"`

	// 映射id
	DataclassId string `json:"dataclass_id"`

	// 数据源
	DataSource string `json:"data_source"`

	// 描述信息
	Description string `json:"description"`

	Classifier *DpeClassifyCreate `json:"classifier"`

	Mapper *CreateDpeMappingRequestBody `json:"mapper,omitempty"`
}

func (o CreateDpeClassifyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDpeClassifyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDpeClassifyRequestBody", string(data)}, " ")
}
