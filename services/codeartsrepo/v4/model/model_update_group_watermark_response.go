package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupWatermarkResponse Response Object
type UpdateGroupWatermarkResponse struct {

	// **参数解释：** 水印开启状态。 - true，开启水印。 - false，关闭水印。
	Watermark      *bool `json:"watermark,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateGroupWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupWatermarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupWatermarkResponse", string(data)}, " ")
}
