package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseError API响应基类，用于老接口200响应，对文档不呈现。
type BaseError struct {
}

func (o BaseError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseError struct{}"
	}

	return strings.Join([]string{"BaseError", string(data)}, " ")
}
