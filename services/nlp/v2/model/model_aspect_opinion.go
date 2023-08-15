package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AspectOpinion 属性观点列表
type AspectOpinion struct {

	// 属性类别 手机领域：['整体','性价比', '赠品','分期', '配件', '活动', '品牌', '物流派送', '包装', '游戏性能', '系统性能', '芯片', '屏幕', '电池', '自拍', '拍照', '音质', '散热', '防水', '信号', '解锁', '外形设计', '握持手感', '质感', '颜色', '内存/容量', '客服/售后', '其他']
	AspectCategory string `json:"aspect_category"`

	// 属性词，与对应的描述词至少出现其中之一，可能为null。
	AspectTerm string `json:"aspect_term"`

	// 描述词，与对应的属性词至少出现其中之一，可能为null。
	OpinionTerm string `json:"opinion_term"`

	// 共4个数字，分别表示属性词和描述词在文本中的起始位置和结束位置。若属性词为null，则1, 2两位不展示；若描述词为null，则3, 4位不展示。
	Span []int32 `json:"span"`

	// 情感标签，0：负向，1：正向
	Label int32 `json:"label"`

	// 情感标签置信度
	Confidence *float32 `json:"confidence,omitempty"`

	// 属性-描述词片段所对应的标签。若分类为'其他'，则不给出标签，返回null。
	Tag *string `json:"tag,omitempty"`
}

func (o AspectOpinion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AspectOpinion struct{}"
	}

	return strings.Join([]string{"AspectOpinion", string(data)}, " ")
}
