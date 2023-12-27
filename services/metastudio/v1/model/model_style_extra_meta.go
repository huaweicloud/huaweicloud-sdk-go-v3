package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StyleExtraMeta 风格额外信息
type StyleExtraMeta struct {

	// 是否支持照片建模
	PictureModelingEnable *bool `json:"picture_modeling_enable,omitempty"`

	// 是否支持模型编辑
	EditEnable *bool `json:"edit_enable,omitempty"`

	// 编辑使用引擎
	EditEngine *string `json:"edit_engine,omitempty"`

	// 照片建模算法调用的模型类型
	ModelId *string `json:"model_id,omitempty"`
}

func (o StyleExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMeta struct{}"
	}

	return strings.Join([]string{"StyleExtraMeta", string(data)}, " ")
}
