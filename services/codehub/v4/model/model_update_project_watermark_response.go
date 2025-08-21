package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectWatermarkResponse Response Object
type UpdateProjectWatermarkResponse struct {

	// **参数解释：** 水印开启状态。 - true，开启水印。 - false，关闭水印。
	Watermark      *bool `json:"watermark,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateProjectWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectWatermarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectWatermarkResponse", string(data)}, " ")
}
