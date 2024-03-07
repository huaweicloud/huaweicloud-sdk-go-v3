package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPreviewResponse Response Object
type CreateAccessPreviewResponse struct {

	// 访问预览的唯一标识符。
	AccessPreviewId *string `json:"access_preview_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateAccessPreviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPreviewResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessPreviewResponse", string(data)}, " ")
}
