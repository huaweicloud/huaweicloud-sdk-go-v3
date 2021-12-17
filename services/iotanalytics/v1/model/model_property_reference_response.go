package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性引用
type PropertyReferenceResponse struct {
	// 属性引用类型，引用本资产属性（this）、引用其他资产属性（single）、引用子资产属性（children）

	Type string `json:"type"`
	// 引用属性所属的资产模型ID，该字段仅当type为“引用其他资产属性”或“引用子资产属性”时有效；使用导入模型和导出模型接口时，该字段无效

	AssetModelId *string `json:"asset_model_id,omitempty"`
	// 引用属性所属的资产模型名称，请求中携带该字段时可以不携带asset_model_id字段

	AssetModelName *string `json:"asset_model_name,omitempty"`
	// 引用属性的名称

	PropertyName string `json:"property_name"`
	// 引用的资产ID，修改资产时携带null表示置空

	AssetId *string `json:"asset_id,omitempty"`
}

func (o PropertyReferenceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyReferenceResponse struct{}"
	}

	return strings.Join([]string{"PropertyReferenceResponse", string(data)}, " ")
}
