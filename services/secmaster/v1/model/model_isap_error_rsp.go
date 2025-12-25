package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapErrorRsp 失败时返回的错误对象
type IsapErrorRsp struct {

	// 错误码
	ErrorCode string `json:"error_code"`

	// 错误描述
	ErrorMsg string `json:"error_msg"`
}

func (o IsapErrorRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapErrorRsp struct{}"
	}

	return strings.Join([]string{"IsapErrorRsp", string(data)}, " ")
}
