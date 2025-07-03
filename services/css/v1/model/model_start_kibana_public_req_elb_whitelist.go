package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartKibanaPublicReqElbWhitelist struct {

	// 是否开启白名单。 - true: 开启白名单。 - false: 关闭白名单。
	EnableWhiteList bool `json:"enable_white_list"`

	// 白名单。
	WhiteList string `json:"white_list"`
}

func (o StartKibanaPublicReqElbWhitelist) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKibanaPublicReqElbWhitelist struct{}"
	}

	return strings.Join([]string{"StartKibanaPublicReqElbWhitelist", string(data)}, " ")
}
