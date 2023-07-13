package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetModelResponse Response Object
type UpdateAssetModelResponse struct {

	// 模型ID
	AssetModelId *string `json:"asset_model_id,omitempty"`

	// 模型名称
	Name *string `json:"name,omitempty"`

	// 模型显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 属性集
	Properties *[]PropertyModelResponse `json:"properties,omitempty"`

	// 分析任务集
	Analyses *[]AnalysisModelResponse `json:"analyses,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 修改时间
	ModifiedTime   *string `json:"modified_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAssetModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetModelResponse", string(data)}, " ")
}
