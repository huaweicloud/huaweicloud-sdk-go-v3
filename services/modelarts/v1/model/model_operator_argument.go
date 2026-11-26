package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperatorArgument struct {

	// **参数解释**：参数键
	Key string `json:"key"`

	// **参数解释**：参数名称
	Name string `json:"name"`

	Type *OperatorArgTypeEnum `json:"type"`

	// **参数解释**：参数说明
	Tips *string `json:"tips,omitempty"`

	// **参数解释**：参数最小值，数值型参数必填
	Min *float32 `json:"min,omitempty"`

	// **参数解释**：参数最大值，数值型参数必填
	Max *float32 `json:"max,omitempty"`

	// **参数解释**：是否为区间型数值
	Between *bool `json:"between,omitempty"`

	// **参数解释**：参数选项列表
	Items *[]OperatorArgItem `json:"items,omitempty"`

	// **参数解释**：参数默认值，用户不可见且必填的参数必须提供默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**：参数是否必填，默认false
	Required *bool `json:"required,omitempty"`

	// **参数解释**：该参数界面是否展示给用户配置，默认true
	Visible *bool `json:"visible,omitempty"`

	// **参数解释**：参数前端组件类型
	ArgDisplayType *string `json:"arg_display_type,omitempty"`

	// **参数解释**：数据集支持的任务类型，type为dataset时可用
	TaskCategory *[]string `json:"task_category,omitempty"`

	// **参数解释**：数据集来源，type为dataset时可用
	DatasetCatalog *[]string `json:"dataset_catalog,omitempty"`

	// **参数解释**：资产类型,type为large_model时可用,多个以','分隔
	AssetTypes *string `json:"asset_types,omitempty"`

	// **参数解释**：资产代码,type为large_model时可用,多个以','分隔
	AssetCodes *string `json:"asset_codes,omitempty"`

	// **参数解释**：模型的子资产类型,type为large_model时可用
	SubAssetType *string `json:"sub_asset_type,omitempty"`

	// **参数解释**：当前参数所依赖的其他参数信息列表
	DependArgs *[]DependArg `json:"depend_args,omitempty"`
}

func (o OperatorArgument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperatorArgument struct{}"
	}

	return strings.Join([]string{"OperatorArgument", string(data)}, " ")
}
