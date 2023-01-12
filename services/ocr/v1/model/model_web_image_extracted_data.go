package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type WebImageExtractedData struct {
	ContactInfo *WebImageContactInfo `json:"contact_info,omitempty"`

	ImageSize *WebImageImageSize `json:"image_size,omitempty"`
}

func (o WebImageExtractedData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageExtractedData struct{}"
	}

	return strings.Join([]string{"WebImageExtractedData", string(data)}, " ")
}
