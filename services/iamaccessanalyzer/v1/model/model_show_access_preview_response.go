package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessPreviewResponse Response Object
type ShowAccessPreviewResponse struct {
	AccessPreview  *AccessPreview `json:"access_preview,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAccessPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessPreviewResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessPreviewResponse", string(data)}, " ")
}
