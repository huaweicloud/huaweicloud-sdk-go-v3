package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkMetadataLabels 网络资源的标签信息。
type NetworkMetadataLabels struct {

	// **参数解释**：资源池的显示名称。 **取值范围**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为4-32。
	OsModelartsName string `json:"os.modelarts/name"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc) **取值范围**：不涉及。
	OsModelartsWorkspaceId *string `json:"os.modelarts/workspace.id,omitempty"`
}

func (o NetworkMetadataLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkMetadataLabels struct{}"
	}

	return strings.Join([]string{"NetworkMetadataLabels", string(data)}, " ")
}
