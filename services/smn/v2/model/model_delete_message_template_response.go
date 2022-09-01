package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteMessageTemplateResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMessageTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteMessageTemplateResponse", string(data)}, " ")
}
