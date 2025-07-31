package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerExtraLinuxInfo 沙箱类型,linux沙箱专用
type ContainerExtraLinuxInfo struct {

	// 操作系统: - ubt : ubuntu - centos : centos - debian, - redhat, - opensuse, - kylin - uos - euleros
	Os *string `json:"os,omitempty"`
}

func (o ContainerExtraLinuxInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerExtraLinuxInfo struct{}"
	}

	return strings.Join([]string{"ContainerExtraLinuxInfo", string(data)}, " ")
}
