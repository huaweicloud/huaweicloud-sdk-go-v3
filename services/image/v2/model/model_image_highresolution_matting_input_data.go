package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingInputData struct {

	// OBS桶名，当输入为obs类型时必填。
	Bucket *string `json:"bucket,omitempty"`

	// OBS的路径，当输入为obs类型时必填。
	Path *string `json:"path,omitempty"`

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
