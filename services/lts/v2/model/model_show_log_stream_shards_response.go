package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogStreamShardsResponse Response Object
type ShowLogStreamShardsResponse struct {

	// 错误码
	ErrorCode *string `json:"errorCode,omitempty"`

	// 是否完全查询
	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`

	// 查询结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLogStreamShardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogStreamShardsResponse struct{}"
	}

	return strings.Join([]string{"ShowLogStreamShardsResponse", string(data)}, " ")
}
