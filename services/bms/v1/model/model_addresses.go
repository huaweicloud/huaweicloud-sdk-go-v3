package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// addresses字段数据结构说明
type Addresses struct {
	// 裸金属服务器所属网络信息。key表示裸金属服务器使用的虚拟私有云的ID。value为网络详细信息

	VpcId []Address `json:"vpc_id"`
}

func (o Addresses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Addresses struct{}"
	}

	return strings.Join([]string{"Addresses", string(data)}, " ")
}
