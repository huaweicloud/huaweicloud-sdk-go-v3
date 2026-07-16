package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotebookTagsRequest Request Object
type CreateNotebookTagsRequest struct {

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ResourceId string `json:"resource_id"`

	// **参数解释**：工作空间ID。获取方法请参见[[查询工作空间列表](ListWorkspace.xml)](tag:hc,hk)。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Body *CreateTmsTagsRequest `json:"body,omitempty"`
}

func (o CreateNotebookTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotebookTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateNotebookTagsRequest", string(data)}, " ")
}
