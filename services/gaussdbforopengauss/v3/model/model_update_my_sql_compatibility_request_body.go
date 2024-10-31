package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMySqlCompatibilityRequestBody struct {

	// MySQL兼容端口，修改端口的时候必填，不填默认保持原有设定端口，可选范围为：0, 1024-39998。 - 取值为0，标识关闭MySQL兼容端口服务。 - 限制端口： 2378,2379,2380,4999,5000,5999,6000,6001,6500,8000-8006,8097,8098,12016,12017,20049,20050,21731,21732,32122,32123,32124。 - 与数据库对外端口互斥。
	Port *string `json:"port,omitempty"`
}

func (o UpdateMySqlCompatibilityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMySqlCompatibilityRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMySqlCompatibilityRequestBody", string(data)}, " ")
}
