package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Image 镜像信息。
type Image struct {

	// 镜像id。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像运行需要的最小磁盘容量，单位为GB。取值为40～1024GB。
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 创建时间，格式为UTC时间，yyyy-MM-ddTHH:mm:ssZ。
	CreatedAt *string `json:"created_at,omitempty"`

	// 镜像文件的大小，单位为字节。
	ImageSize *string `json:"image_size,omitempty"`

	// 更新时间，格式为UTC时间，yyyy-MM-ddTHH:mm:ssZ。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 市场镜像的产品ID。
	ProductCode *string `json:"product_code,omitempty"`

	// 镜像支持的最大内存，单位为MB。取值可以参考云服务器规格限制，一般不设置。
	MaxRam *string `json:"max_ram,omitempty"`

	// 云市场资源类型编码。
	MarketResourceTypeCode *string `json:"market_resource_type_code,omitempty"`

	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，默认设置为0。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 操作系统类型，目前取值Linux， Windows，Other。
	OsType *string `json:"os_type,omitempty"`

	// 镜像类型，目前支持以下类型： 公共镜像：gold 私有镜像：private 共享镜像：shared。
	ImageType *string `json:"image_type,omitempty"`

	// 镜像标签列表。
	Tags *[]string `json:"tags,omitempty"`

	// 镜像平台分类。
	Platform *string `json:"platform,omitempty"`

	// 操作系统位数，一般取值为“32”或者“64”。
	OsBit *string `json:"os_bit,omitempty"`

	// 操作系统具体版本。
	OsVersion *string `json:"os_version,omitempty"`

	// 镜像名称。
	Name *string `json:"name,omitempty"`

	// 云市场服务类型编码。
	MarketServiceTypeCode *string `json:"market_service_type_code,omitempty"`

	// 镜像状态。
	Status *string `json:"status,omitempty"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}
