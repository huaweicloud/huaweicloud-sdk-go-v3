package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunTextModerationResponse struct {
	Result         *TextDetectionResponseResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                          `json:"-"`
}

func (o RunTextModerationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextModerationResponse struct{}"
	}

	return strings.Join([]string{"RunTextModerationResponse", string(data)}, " ")
}
