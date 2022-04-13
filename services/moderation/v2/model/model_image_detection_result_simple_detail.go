package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 反黄，暴恐检测详情
type ImageDetectionResultSimpleDetail struct {
	// 置信度，取值范围 0-1。

	Confidence *float32 `json:"confidence,omitempty"`
	// 每个检测结果的标签化说明： terrorism：label为对应的涉政暴恐元素信息。 - 涉政暴恐场景当前支持label列表如下：   - normal：正常   - knife：刀   - gun：枪   - fire：火灾   - bloody ：血腥   - terrorist：暴恐组织及标志   - fascist：法西斯组织及标志   - cult：邪教组织及标志   - negative_politics ：涉政负面组织及标志   - negative_political_events：涉政负面事件及标志   - special_characters ：特殊文字   - kidnap：绑架   - corpse：尸体   - riot：暴乱事件   - parade ：游行示威   - sensitive_landmarks：敏感地标   - military_weapon：军事武器   - army：警察部队   - positive_politics：涉政正面组织及标志   - crowd：人群聚集 porn：label为对应的涉黄分类（涉黄、性感等）信息。 - 鉴黄场景当前支持label列表如下：   - normal：正常   - porn：色情   - sexy：性感

	Label *string `json:"label,omitempty"`
}

func (o ImageDetectionResultSimpleDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultSimpleDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultSimpleDetail", string(data)}, " ")
}
