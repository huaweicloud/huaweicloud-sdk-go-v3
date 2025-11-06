package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryWatermarkResponse Response Object
type UpdateRepositoryWatermarkResponse struct {

	// **参数解释：** 水印开启状态。 - true，开启水印。 - false，关闭水印。
	Watermark      *bool `json:"watermark,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateRepositoryWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryWatermarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryWatermarkResponse", string(data)}, " ")
}
