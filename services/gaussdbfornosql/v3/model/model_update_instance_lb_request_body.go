package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceLbRequestBody struct {

	// 新的负载均衡IP。
	Ip string `json:"ip"`
}

func (o UpdateInstanceLbRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceLbRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceLbRequestBody", string(data)}, " ")
}
