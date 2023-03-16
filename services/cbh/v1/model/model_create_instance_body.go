package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建堡垒机实例请求参数。
type CreateInstanceBody struct {
	Server *CbhInstances `json:"server"`
}

func (o CreateInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceBody", string(data)}, " ")
}
