package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModelDto struct {

	// 模型名称
	Name *string `json:"name,omitempty"`

	// 模型ID
	Id *string `json:"id,omitempty"`

	// 模型类型
	Type *string `json:"type,omitempty"`

	// 模型创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 模型结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 创建模型的用户名称
	Creator *string `json:"creator,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 是否打开组织共享
	Shareable *bool `json:"shareable,omitempty"`

	// 模型数据量
	DataQuantity *int32 `json:"data_quantity,omitempty"`

	File *ModelFile `json:"file,omitempty"`

	ValueRange *ValueRange2 `json:"value_range,omitempty"`

	// 模型描述信息
	Description *string `json:"description,omitempty"`

	// 失败提示，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// 模型训练loss信息
	Losses *[]float32 `json:"losses,omitempty"`

	// 模型评估指标
	Metrics *[]ModelMetric `json:"metrics,omitempty"`

	// 基模型id
	BaseModelId *string `json:"base_model_id,omitempty"`

	// 基模型名称
	BaseModelName *string `json:"base_model_name,omitempty"`
}

func (o ModelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelDto struct{}"
	}

	return strings.Join([]string{"ModelDto", string(data)}, " ")
}
