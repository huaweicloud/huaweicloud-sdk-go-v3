package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteResourceConfigurationReqBody struct {

	// 待删除的资源列表。
	Resources []string `json:"resources"`
}

func (o DeleteResourceConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"DeleteResourceConfigurationReqBody", string(data)}, " ")
}
