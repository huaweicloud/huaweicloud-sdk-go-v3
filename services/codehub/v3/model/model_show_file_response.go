package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileResponse Response Object
type ShowFileResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *FileContentInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileResponse struct{}"
	}

	return strings.Join([]string{"ShowFileResponse", string(data)}, " ")
}
