package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseError API响应基类，用于老接口200响应，对文档不呈现
type BaseError struct {

	// 错误码，失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BaseError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseError struct{}"
	}

	return strings.Join([]string{"BaseError", string(data)}, " ")
}
