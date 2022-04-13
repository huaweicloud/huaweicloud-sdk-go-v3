package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改设备设备配置结构体。
type UpdateDesireds struct {
	// 设备配置，内容由产品的$config服务定义。

	Config *interface{} `json:"config,omitempty"`
}

func (o UpdateDesireds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesireds struct{}"
	}

	return strings.Join([]string{"UpdateDesireds", string(data)}, " ")
}
