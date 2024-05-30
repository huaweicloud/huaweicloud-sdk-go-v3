package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageResponse Response Object
type DeleteImageResponse struct {

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageResponse", string(data)}, " ")
}
