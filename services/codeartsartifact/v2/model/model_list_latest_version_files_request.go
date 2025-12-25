package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLatestVersionFilesRequest Request Object
type ListLatestVersionFilesRequest struct {

	// **参数解释**： 项目ID，对应\"需求管理 CodeArts Req\"项目唯一标识，私有依赖库首页地址栏url https://{host}/cloudartifact/project/{project_id}/repository中project_id变量的值。 **约束限制**： 字符串长度32。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： name。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： sort_by。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释**： sort_dir。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： limit。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLatestVersionFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatestVersionFilesRequest struct{}"
	}

	return strings.Join([]string{"ListLatestVersionFilesRequest", string(data)}, " ")
}
