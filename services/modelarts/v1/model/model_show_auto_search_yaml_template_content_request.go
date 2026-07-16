package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchYamlTemplateContentRequest Request Object
type ShowAutoSearchYamlTemplateContentRequest struct {

	// 搜索算法类型。
	AlgorithmType string `json:"algorithm_type"`

	// 搜索算法名称。
	AlgorithmName string `json:"algorithm_name"`
}

func (o ShowAutoSearchYamlTemplateContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchYamlTemplateContentRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchYamlTemplateContentRequest", string(data)}, " ")
}
