package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRunningDetailsResponse Response Object
type ShowTaskRunningDetailsResponse struct {

	// **参数解释**： 演化任务启动时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	MetaStartAt *int64 `json:"meta_start_at,omitempty"`

	// **参数解释**： 演化任务完成时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	MetaFinishAt   *int64 `json:"meta_finish_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskRunningDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRunningDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskRunningDetailsResponse", string(data)}, " ")
}
