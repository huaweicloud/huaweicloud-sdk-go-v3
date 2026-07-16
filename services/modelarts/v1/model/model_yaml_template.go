package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// YamlTemplate 自动化搜索作业yaml模板。
type YamlTemplate struct {

	// AutoSearch算法类型，英文描述。
	AlgorithmTypeEn *string `json:"algorithm_type_en,omitempty"`

	// AutoSearch算法类型[，中文描述](tag:hc,hk)。
	AlgorithmTypeZh *string `json:"algorithm_type_zh,omitempty"`

	// 该算法类型下所有算法的名称。
	AlgorithmNames *[]string `json:"algorithm_names,omitempty"`
}

func (o YamlTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "YamlTemplate struct{}"
	}

	return strings.Join([]string{"YamlTemplate", string(data)}, " ")
}
