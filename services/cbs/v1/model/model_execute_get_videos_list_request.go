package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetVideosListRequest Request Object
type ExecuteGetVideosListRequest struct {

	// 范围：大于等于1 默认：10
	Limit int32 `json:"limit"`

	// 长度小于63
	Name string `json:"name"`

	// 范围：大于等于0 默认：0
	Offset int32 `json:"offset"`
}

func (o ExecuteGetVideosListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetVideosListRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetVideosListRequest", string(data)}, " ")
}
