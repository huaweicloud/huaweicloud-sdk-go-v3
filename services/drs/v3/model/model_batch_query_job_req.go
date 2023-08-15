package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryJobReq 查询任务请求体
type BatchQueryJobReq struct {

	// 查询任务请求体
	Jobs []string `json:"jobs"`
}

func (o BatchQueryJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryJobReq struct{}"
	}

	return strings.Join([]string{"BatchQueryJobReq", string(data)}, " ")
}
