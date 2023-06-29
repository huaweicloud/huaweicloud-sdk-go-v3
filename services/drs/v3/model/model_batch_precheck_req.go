package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPrecheckReq 批量预检查请求体
type BatchPrecheckReq struct {

	// 批量预检查请求列表
	Jobs []PreCheckInfo `json:"jobs"`
}

func (o BatchPrecheckReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPrecheckReq struct{}"
	}

	return strings.Join([]string{"BatchPrecheckReq", string(data)}, " ")
}
