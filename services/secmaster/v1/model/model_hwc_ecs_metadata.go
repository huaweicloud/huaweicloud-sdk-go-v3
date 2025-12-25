package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsMetadata 弹性云服务器元数据。
type HwcEcsMetadata struct {

	// 云服务器操作系统对应的镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像类型，目前支持： 公共镜像（gold） 私有镜像（private） 共享镜像（shared）
	ImageType *string `json:"image_type,omitempty"`

	// 云服务器操作系统对应的镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// 操作系统位数，一般取值为“32”或者“64”。
	OsBit *string `json:"os_bit,omitempty"`

	// 操作系统类型，取值为：Linux、Windows。
	OsType *string `json:"os_type,omitempty"`

	// 云服务器所属的虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 云服务器对应的资源规格。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 云服务器对应的资源类型。 取值为“1”，代表资源类型为云服务器。
	ResourceType *string `json:"resource_type,omitempty"`

	// 委托的名称。 委托是由租户管理员在统一身份认证服务（Identity and Access Management，IAM）上创建的，可以为弹性云服务器提供访问云服务器的临时凭证。
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o HwcEcsMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsMetadata struct{}"
	}

	return strings.Join([]string{"HwcEcsMetadata", string(data)}, " ")
}
