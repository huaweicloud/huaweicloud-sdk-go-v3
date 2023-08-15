package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenGaussShard DN分片扩容时必填
type OpenGaussShard struct {

	// 新增的DN分片扩容数大小
	Count int32 `json:"count"`
}

func (o OpenGaussShard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussShard struct{}"
	}

	return strings.Join([]string{"OpenGaussShard", string(data)}, " ")
}
