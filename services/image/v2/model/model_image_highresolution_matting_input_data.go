package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingInputData struct {

	// url输入源的地址，当输入为url类型时必填。 长度不超过1000。输入的图片大小不能大于20M，长边不能大于10000px。
	Url *string `json:"url,omitempty"`

	// 数据标识。多输入场景下必选，值由算法定义；单输入场景非必选。
	Key *string `json:"key,omitempty"`
}

func (o ImageHighresolutionMattingInputData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingInputData struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingInputData", string(data)}, " ")
}
