package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeType 卷类型对象
type VolumeType struct {

	// 硬盘类型的ID。
	Id *string `json:"id,omitempty"`

	// 硬盘类型名称。
	Name *string `json:"name,omitempty"`

	// 磁盘类型的状态： 取值范围：normal(当前磁盘类型存在未售罄站点)，sellout(当前磁盘类型所有站点均售罄)
	Status *string `json:"status,omitempty"`
}

func (o VolumeType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeType struct{}"
	}

	return strings.Join([]string{"VolumeType", string(data)}, " ")
}
