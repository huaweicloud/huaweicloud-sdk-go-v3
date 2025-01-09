package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Watermark struct {

	// 是否开启水印策略设置。取值为：false：表示关闭。true：表示开启。
	WatermarkEnable *bool `json:"watermark_enable,omitempty"`

	Options *WatermarkOptions `json:"options,omitempty"`
}

func (o Watermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Watermark struct{}"
	}

	return strings.Join([]string{"Watermark", string(data)}, " ")
}
