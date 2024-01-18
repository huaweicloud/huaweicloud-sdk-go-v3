package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDialogUrlResponse Response Object
type CreateDialogUrlResponse struct {

	// 对话链接。
	Url *string `json:"url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDialogUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDialogUrlResponse struct{}"
	}

	return strings.Join([]string{"CreateDialogUrlResponse", string(data)}, " ")
}
