package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SfsInfo struct {

	// BCS服务下的SFS文件系统名称
	PvcName *string `json:"pvc_name,omitempty" xml:"pvc_name"`

	// BCS服务网络存储名称
	Name *string `json:"name,omitempty" xml:"name"`

	// BCS服务网络存储地址
	Addr *string `json:"addr,omitempty" xml:"addr"`

	// BCS服务网络存储类型
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o SfsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SfsInfo struct{}"
	}

	return strings.Join([]string{"SfsInfo", string(data)}, " ")
}
