package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性观点列表
type AspectAdvanceOpinion struct {

	// 属性类别 手机领域：['整体','内存','外形设计','屏幕','性价比','拍照','散热','电池','人脸识别','信号','指纹识别','音质','握持手感','活动配件赠品','防水','客服','物流派送','包装'] 汽车领域：['动力','外观','内饰','空间','操控', '舒适性', '性价比','能耗']
	AspectCategory string `json:"aspect_category"`

	// 情感标签，0：负向，1：正向
	Label int32 `json:"label"`

	// 情感标签置信度
	Confidence *float32 `json:"confidence,omitempty"`

	// 属性描述词，预留参数，暂不支持。
	AspectTerm string `json:"aspect_term"`

	// 观点描述词，预留参数，暂不支持。
	OpinionTerm string `json:"opinion_term"`

	// 属性评价起始位置，预留参数，暂不支持。
	Span []int32 `json:"span"`

	// 观点标签，预留参数，暂不支持。
	Tag *string `json:"tag,omitempty"`
}

func (o AspectAdvanceOpinion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AspectAdvanceOpinion struct{}"
	}

	return strings.Join([]string{"AspectAdvanceOpinion", string(data)}, " ")
}
