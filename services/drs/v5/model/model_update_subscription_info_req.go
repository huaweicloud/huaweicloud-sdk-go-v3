package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubscriptionInfoReq struct {

	// 任务名称 约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。 - 最小长度：4 - 最大长度：50
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 消费时间点，在修改完消费时间点后，拉取到的增量数据从修改后的消费时间点开始。 约束：修改的时间点必须在订阅任务的时间范围内（从任务创建到当前时间之间），取值为时间戳，例如：1769393264000。
	ConsumeTime *int64 `json:"consume_time,omitempty"`
}

func (o UpdateSubscriptionInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionInfoReq struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionInfoReq", string(data)}, " ")
}
