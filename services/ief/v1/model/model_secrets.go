package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境变量引用密钥时使用。使用ValueFrom时，secret与configmap必须二选一。
type Secrets struct {

	// 密钥的名称
	Name string `json:"name"`

	// 密钥的属性名
	Key string `json:"key"`
}

func (o Secrets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secrets struct{}"
	}

	return strings.Join([]string{"Secrets", string(data)}, " ")
}
