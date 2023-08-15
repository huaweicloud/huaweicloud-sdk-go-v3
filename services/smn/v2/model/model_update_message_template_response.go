package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageTemplateResponse Response Object
type UpdateMessageTemplateResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMessageTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateMessageTemplateResponse", string(data)}, " ")
}
