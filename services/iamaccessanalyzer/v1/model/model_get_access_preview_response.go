package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccessPreviewResponse Response Object
type GetAccessPreviewResponse struct {
	AccessPreview  *AccessPreview `json:"access_preview,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o GetAccessPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccessPreviewResponse struct{}"
	}

	return strings.Join([]string{"GetAccessPreviewResponse", string(data)}, " ")
}
