package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTestConnectionReq 批量测试连接任务请求体
type BatchTestConnectionReq struct {

	// 批量测试连接请求列表。
	Jobs []TestEndPoint `json:"jobs"`
}

func (o BatchTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchTestConnectionReq", string(data)}, " ")
}
