package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateRequest 自动创建虚拟机模板
type TemplateRequest struct {

	// 模板名称 仅由中文字符、下划线、短横线、数字、英文大小写字母组成
	Name string `json:"name"`

	// 是否是通用模板，如果模板关联一个任务，则不算通用模板
	IsTemplate *bool `json:"is_template,omitempty"`

	// Region信息
	Region *string `json:"region,omitempty"`

	// 项目ID
	Projectid *string `json:"projectid,omitempty"`

	// 目标端服务器名称
	TargetServerName *string `json:"target_server_name,omitempty"`

	// 可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 磁盘类型 （非枚举数据，来源于EVS服务） 常见类型如：SAS:串行连接SCSI，SSD:固态硬盘，SATA:串口硬盘等。 详细类型请参考EIP服务API文档中“查询单个云硬盘详情”部分，查看响应参数的中volume_type字段描述
	Volumetype *string `json:"volumetype,omitempty"`

	// 虚拟机规格
	Flavor *string `json:"flavor,omitempty"`

	Vpc *VpcObject `json:"vpc,omitempty"`

	// 网卡信息，支持多个网卡，如果是自动创建，只填一个，ID使用“autoCreate”
	Nics *[]Nics `json:"nics,omitempty"`

	// 安全组，支持多个安全组，如果是自动创建，只填一个，ID使用“autoCreate”
	SecurityGroups *[]SgObject `json:"security_groups,omitempty"`

	Publicip *PublicIp `json:"publicip,omitempty"`

	// 磁盘信息
	Disk *[]TemplateDisk `json:"disk,omitempty"`

	// 数据盘磁盘类型 （非枚举数据，来源于EVS服务） 常见类型如：SAS:串行连接SCSI，SSD:固态硬盘，SATA:串口硬盘等。 详细类型请参考EIP服务API文档中“查询单个云硬盘详情”部分，查看响应参数的中volume_type字段描述
	DataVolumeType *string `json:"data_volume_type,omitempty"`

	// 目的端密码
	TargetPassword *string `json:"target_password,omitempty"`

	// 新建目的虚拟机用户选择的镜像版本Id值
	ImageId *string `json:"image_id,omitempty"`
}

func (o TemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRequest struct{}"
	}

	return strings.Join([]string{"TemplateRequest", string(data)}, " ")
}
