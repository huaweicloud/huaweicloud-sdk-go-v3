package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出配置请求
type OutputRequest struct {
	// 输出参数名称,formulas中定义的name

	Name string `json:"name"`
	// 输出资产ID，填写模型中定义的输出模型对应的某资产ID；创建资产时，如果是输出到本资产的模型，且output_static_asset_id和output_dynamic_asset_id都未配置，则后台自动配置output_static_asset_id为本资产ID；修改资产时，如果output_static_asset_id为null则表示置空

	OutputStaticAssetId *string `json:"output_static_asset_id,omitempty"`
	// 输出资产ID，填写公式动态生成资产ID，可根据入参获取资产ID，如：GetAssetId(\"assetmodelName1\",\"staticPropertyName1\",paramA)；修改资产时，如果output_static_asset_id为null则表示置空

	OutputDynamicAssetId *string `json:"output_dynamic_asset_id,omitempty"`
}

func (o OutputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputRequest struct{}"
	}

	return strings.Join([]string{"OutputRequest", string(data)}, " ")
}
