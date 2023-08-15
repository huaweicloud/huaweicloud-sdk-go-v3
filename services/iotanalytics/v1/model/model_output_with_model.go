package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputWithModel 输出映射
type OutputWithModel struct {

	// 输出参数名称,formulas中定义的name
	Name string `json:"name"`

	// 输出模型ID，如果输出到本模型可以不携带；使用导入模型和导出模型接口时，该字段无效
	OutputAssetModelId *string `json:"output_asset_model_id,omitempty"`

	// 输出模型名称，请求中携带该字段时可以不携带output_asset_model_id
	OutputAssetModelName *string `json:"output_asset_model_name,omitempty"`

	// 输出属性名
	OutputProperty string `json:"output_property"`
}

func (o OutputWithModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputWithModel struct{}"
	}

	return strings.Join([]string{"OutputWithModel", string(data)}, " ")
}
