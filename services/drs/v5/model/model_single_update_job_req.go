package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SingleUpdateJobReq 更新单个任务请求体。
type SingleUpdateJobReq struct {
	Job *UpdateJobReq `json:"job"`
}

func (o SingleUpdateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleUpdateJobReq struct{}"
	}

	return strings.Join([]string{"SingleUpdateJobReq", string(data)}, " ")
}
