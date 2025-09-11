package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddHbaConfRequestBody struct {

	// **参数解释**: 需要新增的hba配置信息。 **约束限制**: 不涉及。
	HbaConfs *[]HbaConfOption `json:"hba_confs,omitempty"`
}

func (o AddHbaConfRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHbaConfRequestBody struct{}"
	}

	return strings.Join([]string{"AddHbaConfRequestBody", string(data)}, " ")
}
