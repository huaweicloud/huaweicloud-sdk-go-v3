package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PropertyReferenceModel 属性引用
type PropertyReferenceModel struct {

	// 属性引用类型，引用本资产属性（this）、引用其他资产属性（single）、引用子资产属性（children）
	Type string `json:"type"`

	// 引用属性所属的资产模型ID，该字段仅当type为“引用其他资产属性”或“引用子资产属性”时有效；使用导入模型和导出模型接口时，该字段无效
	AssetModelId *string `json:"asset_model_id,omitempty"`

	// 引用属性所属的资产模型名称，请求中携带该字段时可以不携带asset_model_id字段
	AssetModelName *string `json:"asset_model_name,omitempty"`

	// 引用属性的名称
	PropertyName string `json:"property_name"`
}

func (o PropertyReferenceModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyReferenceModel struct{}"
	}

	return strings.Join([]string{"PropertyReferenceModel", string(data)}, " ")
}
