package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云堡垒机实例网卡信息。
type Nics struct {

	// 子网ID，字母数字下划线连接符组成。
	SubnetId string `json:"subnet_id"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
