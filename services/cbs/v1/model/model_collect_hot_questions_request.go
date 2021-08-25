package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CollectHotQuestionsRequest struct {
	// qabot编号，UUID格式。

	QabotId string `json:"qabot_id"`
	// 查询的起始时间，long，UTC时间，默认值为0。

	StartTime *string `json:"start_time,omitempty"`
	// 查询的结束时间，long，UTC时间，默认值为当前时间的毫秒数。

	EndTime *string `json:"end_time,omitempty"`
	// 热点问题最多显示的个数，默认值为10，取值范围1-20。

	Top *string `json:"top,omitempty"`
	// 热点问题所属的领域。如果指定领域为非空字符串则从指定领域中查询热点问题，否则从所有标准问题中查询热点问题。

	Domain *string `json:"domain,omitempty"`
	// 统计的目标问题类别id。

	DomainId *string `json:"domain_id,omitempty"`
	// true:根据问答对信息展示热点问题（如：热点问题对应的问答对“你好”发生了修改，变成了 “你好啊”，此时热点问题也将返回 “你好啊”；但是如果这个问题对被删除，则“你好”不会被展示在热点问中） false: 不根据问答对信息展示热点问题。

	Exclude *bool `json:"exclude,omitempty"`
}

func (o CollectHotQuestionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CollectHotQuestionsRequest struct{}"
	}

	return strings.Join([]string{"CollectHotQuestionsRequest", string(data)}, " ")
}
