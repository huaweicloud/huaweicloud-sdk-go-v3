package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchImportConfigsRequestBody struct {

	// 南向IA配置项列表
	Configs *[]BatchImportConfigRequestBody `json:"configs,omitempty" xml:"configs"`
}

func (o BatchImportConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportConfigsRequestBody", string(data)}, " ")
}
