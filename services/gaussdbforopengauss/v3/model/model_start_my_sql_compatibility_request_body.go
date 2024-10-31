package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartMySqlCompatibilityRequestBody struct {

	// MySQL兼容端口，可选范围为：1024-39998。 - 限制端口： 2378,2379,2380,4999,5000,5999,6000,6001,6500,8000-8006,8097,8098,12016,12017,20049,20050,21731,21732,32122,32123,32124。 - 与数据库对外端口互斥。
	Port string `json:"port"`
}

func (o StartMySqlCompatibilityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMySqlCompatibilityRequestBody struct{}"
	}

	return strings.Join([]string{"StartMySqlCompatibilityRequestBody", string(data)}, " ")
}
