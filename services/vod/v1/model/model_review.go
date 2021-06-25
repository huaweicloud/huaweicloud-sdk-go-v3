package model

import (
	"encoding/json"

	"strings"
)

// 媒资审核参数
type Review struct {
	// 审核模板ID

	TemplateId string `json:"template_id"`
	// 截图的时间间隔。单位为秒

	Interval *int32 `json:"interval,omitempty"`
	// 进行政治人物检测时的置信度。 1）  未传参时表示不进行此项检测。 2）  传 -1 表示采用默认的置信度

	Politics *int32 `json:"politics,omitempty"`
	// 进行暴恐元素检测时的置信度。 1)  未传参时表示不进行此项检测。 2)  传 -1 表示采用默认的置信度

	Terrorism *int32 `json:"terrorism,omitempty"`
	// 进行涉黄内容检测时的置信度。 1)  未传参时表示不进行此项检测。 2)  传 -1 表示采用默认的置信度

	Porn *int32 `json:"porn,omitempty"`
}

func (o Review) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Review struct{}"
	}

	return strings.Join([]string{"Review", string(data)}, " ")
}
