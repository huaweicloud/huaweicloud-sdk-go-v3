package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAsyncUpdateJobReq 批量更新指定ID异步任务请求体。
type BatchAsyncUpdateJobReq struct {

	// 批量更新指定ID异步任务请求体。
	Jobs []UpdateJobReq `json:"jobs"`
}

func (o BatchAsyncUpdateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAsyncUpdateJobReq struct{}"
	}

	return strings.Join([]string{"BatchAsyncUpdateJobReq", string(data)}, " ")
}
