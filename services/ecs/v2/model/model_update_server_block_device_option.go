package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerBlockDeviceOption
type UpdateServerBlockDeviceOption struct {

	// 云硬盘随实例释放策略。 -  true：云硬盘随实例释放。  -  false：云硬盘不随实例释放。  > 说明 > > 不支持修改包年/包月计费模式的磁盘。 > 不支持修改共享盘。
	DeleteOnTermination bool `json:"delete_on_termination"`
}

func (o UpdateServerBlockDeviceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerBlockDeviceOption struct{}"
	}

	return strings.Join([]string{"UpdateServerBlockDeviceOption", string(data)}, " ")
}
