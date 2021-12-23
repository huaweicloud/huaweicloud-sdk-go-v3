package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 广告检测详情
type ImageDetectionResultAdDetail struct {
	// 置信度，取值范围 0-1。

	Confidence *float32 `json:"confidence,omitempty"`
	// ad：label为对应的广告识别结果信息 - 广告场景当前支持label列表如下：   - normal：正常   - ad：广告 - 图文审核场景当前支持label列表如下：   - normal：正常   - qr_code：二维   - politics：涉政   - porn：涉黄   - ad：广告   - abuse：辱骂   - contraband：违禁品   - 其他自定义黑库名称

	Label *string `json:"label,omitempty"`
	// 图文审核场景命中的文本列表。

	HitContexts *[]string `json:"hit_contexts,omitempty"`
}

func (o ImageDetectionResultAdDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultAdDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultAdDetail", string(data)}, " ")
}
