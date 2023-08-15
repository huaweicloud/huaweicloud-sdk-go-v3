package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAsyncCreateJobReq 批量异步创建任务请求体。
type BatchAsyncCreateJobReq struct {

	// 异步创建任务请求体。
	Jobs []AsyncCreateJobReq `json:"jobs"`
}

func (o BatchAsyncCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAsyncCreateJobReq struct{}"
	}

	return strings.Join([]string{"BatchAsyncCreateJobReq", string(data)}, " ")
}
