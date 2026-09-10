package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelServiceListRequest Request Object
type ShowModelServiceListRequest struct {

	// 推理类型，分为online和edge
	InferType *string `json:"infer_type,omitempty"`

	// 推理服务所关联的模型ID
	AssetId *string `json:"asset_id,omitempty"`

	// 推理服务所关联的模型类型
	AssetType *string `json:"asset_type,omitempty"`

	// 推理服务所关联的模型子类型
	AssetSubType *string `json:"asset_sub_type,omitempty"`

	// 芯片类型
	ChipType *string `json:"chip_type,omitempty"`

	// 部署平台，Modelarts或者CCE
	Platform *string `json:"platform,omitempty"`

	// 服务状态，INIT/DEPLOYING/RUNNING/SUCCEEDED/FAILED/STOPPED/PENDING/CANCELLED/WAITING
	Status *string `json:"status,omitempty"`

	// 使用类型，private表示用户创建的服务，public表示预置服务
	UseType *string `json:"use_type,omitempty"`

	// 模型名称，支持模糊匹配
	ModelName *string `json:"model_name,omitempty"`

	// 服务名称，支持模糊匹配
	ServiceName *string `json:"service_name,omitempty"`

	// 偏移量，取值范围[0,100000000]，默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 返回限制个数，取值范围[1-1000]，默认值100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及 **取值范围**： - DESC：降序。 - ASC：升序。 **默认取值**： DESC
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ShowModelServiceListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelServiceListRequest struct{}"
	}

	return strings.Join([]string{"ShowModelServiceListRequest", string(data)}, " ")
}
