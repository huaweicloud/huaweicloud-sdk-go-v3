package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageParameter struct {

	// 主机路径， 适用于HostPath的存储类型
	Path *string `json:"path,omitempty"`

	// 配置项、密钥或者PVC的名字， 适用于ConfigMap、Secret和PersistentVolumeClaim的存储类型
	Name *string `json:"name,omitempty"`

	// 挂载的权限，十进制格式，如384
	DefaultMode *int32 `json:"defaultMode,omitempty"`

	// 适用于EmptyDir类型的存储。不传参数为默认的磁盘介质，传参为memory则开启内存存储。
	Medium *string `json:"medium,omitempty"`
}

func (o StorageParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageParameter struct{}"
	}

	return strings.Join([]string{"StorageParameter", string(data)}, " ")
}
