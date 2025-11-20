package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceIp struct {

	// 数据中心
	DataCenter *string `json:"data_center,omitempty"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// 实例ip
	Ip *[]string `json:"ip,omitempty"`
}

func (o SourceIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceIp struct{}"
	}

	return strings.Join([]string{"SourceIp", string(data)}, " ")
}
