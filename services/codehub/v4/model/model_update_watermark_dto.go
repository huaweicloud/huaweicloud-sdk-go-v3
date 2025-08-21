package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWatermarkDto struct {

	// **参数解释：** 水印开启状态。 - true，开启水印。 - false，关闭水印。
	Watermark *bool `json:"watermark,omitempty"`
}

func (o UpdateWatermarkDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkDto struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkDto", string(data)}, " ")
}
