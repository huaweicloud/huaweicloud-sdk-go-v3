package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWhiteListRequestBody struct {

	// 白名单列表
	IpList []IpInfo `json:"ip_list"`
}

func (o UpdateWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhiteListRequestBody", string(data)}, " ")
}
