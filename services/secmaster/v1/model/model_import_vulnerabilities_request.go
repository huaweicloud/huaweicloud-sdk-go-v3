package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportVulnerabilitiesRequest Request Object
type ImportVulnerabilitiesRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *ImportVulnerabilitiesRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ImportVulnerabilitiesRequest", string(data)}, " ")
}
