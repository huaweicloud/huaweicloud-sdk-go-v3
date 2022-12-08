package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 产品信息
type ProductInfoBean struct {

	// 资源名称列表
	AllResourceNames *[]string `json:"allResourceNames,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 展示ID
	DisplayId *string `json:"displayId,omitempty"`

	// 产品ID
	ProductId *string `json:"productId,omitempty"`

	// 产品规格描述
	ProductSpecDesc *string `json:"productSpecDesc,omitempty"`

	// 资源名称
	ResourceName *string `json:"resourceName,omitempty"`

	// 产品支持的数据库数量，或者支持obs的扫描量
	ResourceSize *int32 `json:"resourceSize,omitempty"`

	// 资源容量度量标识，枚举值举例如下：15：mbps（购买带宽时使用），17：gb（购买云硬盘时使用），14：个/次
	ResourceSizeMeasureId *int32 `json:"resourceSizeMeasureId,omitempty"`

	// 产品编码
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// 资源类型
	ResourceType *string `json:"resourceType,omitempty"`

	// 已使用系数
	UsageFactor *string `json:"usageFactor,omitempty"`

	// 已使用容量度量标识
	UsageMeasureId *int32 `json:"usageMeasureId,omitempty"`

	// 已使用值
	UsageValue *float64 `json:"usageValue,omitempty"`
}

func (o ProductInfoBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfoBean struct{}"
	}

	return strings.Join([]string{"ProductInfoBean", string(data)}, " ")
}
