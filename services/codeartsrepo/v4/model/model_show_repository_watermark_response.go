package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryWatermarkResponse Response Object
type ShowRepositoryWatermarkResponse struct {

	// **参数解释：** 水印设置状态。 - true，开启水印。 - false，关闭水印。
	Watermark *bool `json:"watermark,omitempty"`

	// **参数解释：** 仓库水印状态。 - true，开启水印。 - false，关闭水印。
	ViewWatermark  *bool `json:"view_watermark,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowRepositoryWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryWatermarkResponse", string(data)}, " ")
}
