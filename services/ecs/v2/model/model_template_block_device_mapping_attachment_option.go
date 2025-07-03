package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateBlockDeviceMappingAttachmentOption struct {

	// 盘在虚拟机上的挂载顺序，0表示启动盘
	BootIndex *int32 `json:"boot_index,omitempty"`

	// 卷是否随实例一同释放 默认系统盘设置为true随实例释放，数据盘设置为false不随实例释放
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o TemplateBlockDeviceMappingAttachmentOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBlockDeviceMappingAttachmentOption struct{}"
	}

	return strings.Join([]string{"TemplateBlockDeviceMappingAttachmentOption", string(data)}, " ")
}
