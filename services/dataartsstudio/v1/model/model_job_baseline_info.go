package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobBaselineInfo struct {

	// **参数解释**： 作业ID。 **取值范围**： 长度为1~64个字符。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 作业名称。如果作业不存在或不在当前工作空间内，则为空。 **取值范围**： 长度为1~1024个字符。
	JobName *string `json:"job_name,omitempty"`

	// **参数解释**： 作业关联的基线信息列表。如果作业未关联基线，则返回空列表。 **取值范围**： 不涉及。
	Baselines *[]BaselineInfo `json:"baselines,omitempty"`
}

func (o JobBaselineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBaselineInfo struct{}"
	}

	return strings.Join([]string{"JobBaselineInfo", string(data)}, " ")
}
