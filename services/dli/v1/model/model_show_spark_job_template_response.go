package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobTemplateResponse Response Object
type ShowSparkJobTemplateResponse struct {

	// 模板类型
	Type *string `json:"type,omitempty"`

	// 模板id
	Id *string `json:"id,omitempty"`

	// 模板名字
	Name *string `json:"name,omitempty"`

	Body *SparkJobTemplateDetail `json:"body,omitempty"`

	// 组名
	Group *string `json:"group,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 模板拥有者
	Owner          *string `json:"owner,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobTemplateResponse", string(data)}, " ")
}
