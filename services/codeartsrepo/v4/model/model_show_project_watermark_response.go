package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectWatermarkResponse Response Object
type ShowProjectWatermarkResponse struct {

	// **参数解释：** 水印设置状态。 - true，开启水印。 - false，关闭水印。
	Watermark *bool `json:"watermark,omitempty"`

	// **参数解释：** 当前用户是否有权限更新水印设置。 - true，有权限更新。 - false，无权限更新。
	CanUpdate      *bool `json:"can_update,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowProjectWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectWatermarkResponse", string(data)}, " ")
}
