package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBucketsResponse Response Object
type AddBucketsResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBucketsResponse struct{}"
	}

	return strings.Join([]string{"AddBucketsResponse", string(data)}, " ")
}
