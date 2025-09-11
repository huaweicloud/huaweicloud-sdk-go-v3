package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindLtsConfigRequestBody struct {

	// **参数解释**: 下发配置的实例日志列表。列表最大长度为1000。 **约束限制**: 不涉及。
	List []InstanceSaveLtsConfigOption `json:"list"`
}

func (o BindLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"BindLtsConfigRequestBody", string(data)}, " ")
}
