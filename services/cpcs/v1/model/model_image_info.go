package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageInfo struct {

	// 镜像ID
	ImageId string `json:"image_id"`

	// 镜像名称
	ImageName string `json:"image_name"`

	// 镜像所属的服务
	ServiceType string `json:"service_type"`

	// 镜像的系统架构： - **X86_64** : X86 - **ARRCH** : ARM
	ArchType string `json:"arch_type"`

	// 规格ID
	SpecificationId string `json:"specification_id"`

	// 镜像的录入时间，UNIX时间戳，单位毫秒
	CreateTime string `json:"create_time"`

	// 版本类型
	VersionType string `json:"version_type"`

	// domain白名单列表，多个之间用逗号分隔
	TrustDomain string `json:"trust_domain"`

	// 厂商名称
	VendorName string `json:"vendor_name"`

	// 厂商版本号
	VendorImageVersion string `json:"vendor_image_version"`

	// 密码服务依赖的(厂商)平台版本号
	CcspVersionNeed string `json:"ccsp_version_need"`

	// 华为版本号
	HwImageVersion string `json:"hw_image_version"`

	// 描述
	Description string `json:"description"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}
