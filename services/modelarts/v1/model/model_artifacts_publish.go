package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArtifactsPublish 训练产物自动发布配置
type ArtifactsPublish struct {

	// 是否是中间产物，false-是模型产物，true-是中间产物
	IsCkpt *bool `json:"is_ckpt,omitempty"`

	// 断点ID,ckpt发布时使用
	ArtifactId *string `json:"artifact_id,omitempty"`

	// 模型产物发布后资产名称，默认{源模型名字}-{训练类型}-{训练时间}
	AssetName *string `json:"asset_name,omitempty"`

	// 全局可见性，用来控制资产是当前空间可见或者全部空间可见，取值current|all。
	Visibility *string `json:"visibility,omitempty"`

	// 发布资产描述信息，{任务名}的最终产出模型
	Description *string `json:"description,omitempty"`

	// 模型发布方式
	PublishAssetType *string `json:"publish_asset_type,omitempty"`

	// 资产来源
	AssetSourceType *string `json:"asset_source_type,omitempty"`

	// 选择模型。
	AssetCode *string `json:"asset_code,omitempty"`

	// 版本号。
	AssetVersion *string `json:"asset_version,omitempty"`

	// 版本描述。
	VersionDescription *string `json:"version_description,omitempty"`
}

func (o ArtifactsPublish) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactsPublish struct{}"
	}

	return strings.Join([]string{"ArtifactsPublish", string(data)}, " ")
}
