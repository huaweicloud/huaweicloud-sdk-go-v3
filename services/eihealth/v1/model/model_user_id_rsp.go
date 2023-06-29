package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserIdRsp 操作用户返回体
type UserIdRsp struct {

	// 用户id
	Id *string `json:"id,omitempty"`
}

func (o UserIdRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserIdRsp struct{}"
	}

	return strings.Join([]string{"UserIdRsp", string(data)}, " ")
}
