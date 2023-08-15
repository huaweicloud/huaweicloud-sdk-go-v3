package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamGraphResponse Response Object
type CreateStreamGraphResponse struct {

	//
	IsSuccess *bool `json:"is_success,omitempty"`

	//
	Message *string `json:"message,omitempty"`

	//
	ErrorCode *string `json:"error_code,omitempty"`

	// 静态流图的描述信息
	StreamGraph    *string `json:"stream_graph,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStreamGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamGraphResponse struct{}"
	}

	return strings.Join([]string{"CreateStreamGraphResponse", string(data)}, " ")
}
