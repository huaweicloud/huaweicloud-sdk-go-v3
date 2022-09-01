package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeWebImageResponse struct {
	Result         *WebImageResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeWebImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeWebImageResponse struct{}"
	}

	return strings.Join([]string{"RecognizeWebImageResponse", string(data)}, " ")
}
