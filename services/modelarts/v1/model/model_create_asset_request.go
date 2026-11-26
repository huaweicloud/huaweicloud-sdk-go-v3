package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetRequest Request Object
type CreateAssetRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**： OBS属主租户token。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XStorageToken *string `json:"X-Storage-Token,omitempty"`

	Body *CreateAssetReq `json:"body,omitempty"`
}

func (o CreateAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetRequest", string(data)}, " ")
}
