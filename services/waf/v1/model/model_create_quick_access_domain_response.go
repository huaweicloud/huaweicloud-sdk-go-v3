package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQuickAccessDomainResponse Response Object
type CreateQuickAccessDomainResponse struct {

	// **参数解释：** 异步任务的状态，如RUNNING（运行中）、SUCCESS（成功）、FAILED（失败） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 异步任务的描述，说明当前任务的进展或结果 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释：** 异步任务的ID，用于查询任务进度或结果的唯一标识 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQuickAccessDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQuickAccessDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateQuickAccessDomainResponse", string(data)}, " ")
}
