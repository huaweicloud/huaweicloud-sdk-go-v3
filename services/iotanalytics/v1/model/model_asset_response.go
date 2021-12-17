package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetResponse struct {
	// 资产ID

	AssetId *string `json:"asset_id,omitempty"`
	// 资产模型ID

	AssetModelId *string `json:"asset_model_id,omitempty"`
	// 资产模型名称

	AssetModelName *string `json:"asset_model_name,omitempty"`
	// 资产名称

	Name *string `json:"name,omitempty"`
	// 资产显示名称

	DisplayName *string `json:"display_name,omitempty"`
	// 属性集

	Properties *[]PropertyResponse `json:"properties,omitempty"`
	// 分析任务集

	Analyses *[]AnalysisResponse `json:"analyses,omitempty"`
	// 根资产ID

	Root *string `json:"root,omitempty"`
	// 父资产ID，根资产的父资产ID为null

	Parent *string `json:"parent,omitempty"`
	// 子资产ID集

	Children *[]string `json:"children,omitempty"`
	// 资产状态，正常状态（ACTIVE），异常状态（INACTIVE）；只有草稿态（SKETCH）资产有此状态；资产处于异常状态的场景有：1、该资产存在未填写设备ID的测量数据类别的属性；2、该资产存在未填写静态值的静态配置类别的属性；3、该资产存在分析任务，该分析任务的输入参数存在属性引用类型为引用其他资产属性，且没有为该输入参数配置引用的其他资产的资产ID

	State *string `json:"state,omitempty"`
	// 资产发布状态，发布中（PUBLISHING），发布完成（PUBLISHED）；只能对草稿态（SKETCH）的根资产进行发布，也只有草稿态的根资产有此字段；如果根资产从未发布过则值为null

	PublishState *string `json:"publish_state,omitempty"`
	// 创建时间，格式\"yyyy-MM-dd'T'HH:mm:ss'Z'\"

	CreatedTime *string `json:"created_time,omitempty"`
	// 修改时间，格式\"yyyy-MM-dd'T'HH:mm:ss'Z'\"

	ModifiedTime *string `json:"modified_time,omitempty"`
	// 发布时间，只能对草稿态（SKETCH）的根资产进行发布，也只有草稿态的根资产有此字段；如果从未发布过则值为null；格式\"yyyy-MM-dd'T'HH:mm:ss'Z'\"

	PublishedTime *string `json:"published_time,omitempty"`
}

func (o AssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetResponse struct{}"
	}

	return strings.Join([]string{"AssetResponse", string(data)}, " ")
}
