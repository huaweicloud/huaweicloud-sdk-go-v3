package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改设备设备配置结构体。
type UpdateDesireds struct {

	// 设备配置，内容由产品的$config服务定义。
	Config *interface{} `json:"config,omitempty" xml:"config"`
}

func (o UpdateDesireds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesireds struct{}"
	}

	return strings.Join([]string{"UpdateDesireds", string(data)}, " ")
}
