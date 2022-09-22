package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionUpdateReq struct {

	// 目标连接描述
	Description *string `json:"description,omitempty"`
}

func (o ConnectionUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionUpdateReq struct{}"
	}

	return strings.Join([]string{"ConnectionUpdateReq", string(data)}, " ")
}
