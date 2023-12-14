package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppReq struct {

	// APP的名称，用户数据消费程序的唯一标识符。  应用名称由字母、数字、下划线和中划线组成，长度为1～200个字符。
	AppName string `json:"app_name"`
}

func (o CreateAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppReq struct{}"
	}

	return strings.Join([]string{"CreateAppReq", string(data)}, " ")
}
