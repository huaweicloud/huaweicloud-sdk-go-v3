package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferProjectReq 转移项目请求体
type TransferProjectReq struct {

	// 转移用户id
	UserId string `json:"user_id"`
}

func (o TransferProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferProjectReq struct{}"
	}

	return strings.Join([]string{"TransferProjectReq", string(data)}, " ")
}
