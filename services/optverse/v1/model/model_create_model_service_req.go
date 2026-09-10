package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelServiceReq 创建模型服务请求体
type CreateModelServiceReq struct {

	// 部署名称
	Name string `json:"name"`

	// 推理类型
	RequestMode string `json:"request_mode"`

	// 描述
	Description *string `json:"description,omitempty"`

	Platform *Platform `json:"platform"`

	// 训练产物OBS地址
	TrainObsOutput *string `json:"train_obs_output,omitempty"`

	// 资产ID
	AssetId string `json:"asset_id"`

	// 对话ID
	ChatId *string `json:"chat_id,omitempty"`

	InferType *InferType `json:"infer_type"`

	ServiceConfig *ModelServiceConfig `json:"service_config"`
}

func (o CreateModelServiceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelServiceReq struct{}"
	}

	return strings.Join([]string{"CreateModelServiceReq", string(data)}, " ")
}
