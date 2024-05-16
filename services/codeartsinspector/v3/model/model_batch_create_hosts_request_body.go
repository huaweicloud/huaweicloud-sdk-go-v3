package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateHostsRequestBody struct {

	// 主机资产列表
	Hosts []HostItem `json:"hosts"`
}

func (o BatchCreateHostsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateHostsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateHostsRequestBody", string(data)}, " ")
}
