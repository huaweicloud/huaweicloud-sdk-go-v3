package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDisasterRecordRequestBody struct {

	// **参数解释**: 容灾任务ID。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、中划线、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	Id string `json:"id"`
}

func (o DeleteDisasterRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecordRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecordRequestBody", string(data)}, " ")
}
