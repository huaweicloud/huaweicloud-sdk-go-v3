package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAssetModelResponse struct {

	// 模型ID
	AssetModelId *string `json:"asset_model_id,omitempty" xml:"asset_model_id"`

	// 模型名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 模型显示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 属性集
	Properties *[]PropertyModelResponse `json:"properties,omitempty" xml:"properties"`

	// 分析任务集
	Analyses *[]AnalysisModelResponse `json:"analyses,omitempty" xml:"analyses"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间
	ModifiedTime   *string `json:"modified_time,omitempty" xml:"modified_time"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetModelResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetModelResponse", string(data)}, " ")
}
