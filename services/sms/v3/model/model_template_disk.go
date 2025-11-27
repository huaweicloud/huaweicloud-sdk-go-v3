package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateDisk 磁盘模板
type TemplateDisk struct {

	// 磁盘ID
	Id *int64 `json:"id,omitempty"`

	// 磁盘序号，从0开始
	Index int32 `json:"index"`

	// 磁盘名称
	Name string `json:"name"`

	// 磁盘类型，同volumetype字段（长文本信息，非枚举数据，来源于EVS服务） 详细类型请参考EIP服务API文档中“查询单个云硬盘详情”部分，查看响应参数的中volume_type字段描述
	Disktype string `json:"disktype"`

	// 磁盘大小，单位：GB
	Size int64 `json:"size"`

	// 磁盘使用
	DeviceUse *string `json:"device_use,omitempty"`
}

func (o TemplateDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateDisk struct{}"
	}

	return strings.Join([]string{"TemplateDisk", string(data)}, " ")
}
