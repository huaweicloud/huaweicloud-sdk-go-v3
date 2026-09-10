package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelServiceRsp 模型服务详情
type ModelServiceRsp struct {

	// 服务ID
	ServiceId string `json:"service_id"`

	// 服务名称
	ServiceName string `json:"service_name"`

	// 服务描述
	ServiceDesc *string `json:"service_desc,omitempty"`

	// 服务状态，INIT/RUNNING/STOPPED/FAILED
	Status string `json:"status"`

	// 推理类型，分为online和edge
	InferType string `json:"infer_type"`

	// 设备类型
	DeviceType *string `json:"device_type,omitempty"`

	// 芯片类型
	ChipType *string `json:"chip_type,omitempty"`

	// 请求类型
	RequestMode *string `json:"request_mode,omitempty"`

	// 推理服务所关联的模型ID
	AssetId string `json:"asset_id"`

	// 推理服务所关联的模型名称
	AssetName *string `json:"asset_name,omitempty"`

	// 推理服务所关联的模型类型
	AssetType *string `json:"asset_type,omitempty"`

	// 推理服务所关联的模型子类型
	AssetSubType *string `json:"asset_sub_type,omitempty"`

	// API调用地址
	ApiUrl *string `json:"api_url,omitempty"`

	// 关联对话ID
	ChatId *string `json:"chat_id,omitempty"`

	// 模型部署服务的用户ID
	UserId *string `json:"user_id,omitempty"`

	// 模型部署服务的用户名称
	UserName *string `json:"user_name,omitempty"`

	// 部署平台，Modelarts或者CCE
	Platform string `json:"platform"`

	// 训练产物OBS地址
	TrainObsOutput *string `json:"train_obs_output,omitempty"`

	// 使用类型，private表示用户创建的服务，public表示预置服务
	UseType *string `json:"use_type,omitempty"`

	ServiceConfig *ModelServiceConfig `json:"service_config,omitempty"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

func (o ModelServiceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelServiceRsp struct{}"
	}

	return strings.Join([]string{"ModelServiceRsp", string(data)}, " ")
}
