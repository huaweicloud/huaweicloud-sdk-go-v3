package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusVoIpd 工作项状态查询接口返回状态数据
type StatusVoIpd struct {

	// **参数解释**： 状态名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 工作项的状态属性。 **取值范围**： START、IN_PROGRESS、END。
	Belonging *string `json:"belonging,omitempty"`
}

func (o StatusVoIpd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusVoIpd struct{}"
	}

	return strings.Join([]string{"StatusVoIpd", string(data)}, " ")
}
