package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseModel 基模型
type BaseModel struct {

	// 类型请从ListBaseModels列举基模型接口响应中获取
	BaseModelType string `json:"base_model_type"`

	// 支持的最大token数
	MaxTokens int32 `json:"max_tokens"`

	// 消耗MU数量
	NumOfMu int32 `json:"num_of_mu"`
}

func (o BaseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseModel struct{}"
	}

	return strings.Join([]string{"BaseModel", string(data)}, " ")
}
