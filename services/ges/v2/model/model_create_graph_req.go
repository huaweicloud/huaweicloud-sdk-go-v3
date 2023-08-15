package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphReq 创建图请求体
type CreateGraphReq struct {
	Graph *CreateGraphReqGraph `json:"graph"`
}

func (o CreateGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReq struct{}"
	}

	return strings.Join([]string{"CreateGraphReq", string(data)}, " ")
}
