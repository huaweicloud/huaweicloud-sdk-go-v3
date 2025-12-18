package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationUpdateReqV3 struct {

	// **参数解释**：  修改参数组相关信息。  **参数范围**：  不涉及。
	UpdatePara *interface{} `json:"update_para,omitempty"`
}

func (o ConfigurationUpdateReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationUpdateReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationUpdateReqV3", string(data)}, " ")
}
