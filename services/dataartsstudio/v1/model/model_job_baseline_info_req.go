package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobBaselineInfoReq struct {

	// **参数解释**： 待查询的作业ID列表。 **约束限制**： 必须传入。 **取值范围**： 列表元素个数为1~100个，每个元素长度为1~64个字符。 **默认取值**： 不涉及。
	JobIds []string `json:"job_ids"`
}

func (o JobBaselineInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBaselineInfoReq struct{}"
	}

	return strings.Join([]string{"JobBaselineInfoReq", string(data)}, " ")
}
