package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulnerabilitiesRequest Request Object
type ListVulnerabilitiesRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *VulnerabilityDataObjectSearch `json:"body,omitempty"`
}

func (o ListVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesRequest", string(data)}, " ")
}
