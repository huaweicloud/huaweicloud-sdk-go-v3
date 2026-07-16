package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerRequest **参数解释**：更新DevServer名称的请求体。
type UpdateServerRequest struct {

	// **参数解释**：服务器名称。 **约束限制**：不涉及。 **取值范围**：长度为1至64个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	Name string `json:"name"`
}

func (o UpdateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerRequest", string(data)}, " ")
}
