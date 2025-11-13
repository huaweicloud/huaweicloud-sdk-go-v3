package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateResponseInfo struct {

	// 订阅ID。
	Id *string `json:"id,omitempty"`

	// 执行结果。success，failure
	Result *string `json:"result,omitempty"`

	// 失败报错。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o BatchOperateResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateResponseInfo struct{}"
	}

	return strings.Join([]string{"BatchOperateResponseInfo", string(data)}, " ")
}
