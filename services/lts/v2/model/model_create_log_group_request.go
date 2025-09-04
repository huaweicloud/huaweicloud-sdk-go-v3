package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogGroupRequest Request Object
type CreateLogGroupRequest struct {

	// **参数解释：** 企业项目ID。获取方式请参见：[获取企业项目ID]。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** default。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateLogGroupParams `json:"body,omitempty"`
}

func (o CreateLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLogGroupRequest", string(data)}, " ")
}
