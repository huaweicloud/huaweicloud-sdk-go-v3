package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProject struct {

	// |参数名称：企业项目id| |参数约束及描述：最大长度256|
	Id *string `json:"id,omitempty"`

	// |参数名称：企业项目名称| |参数约束及描述：最大长度256|
	Name *string `json:"name,omitempty"`
}

func (o EnterpriseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProject struct{}"
	}

	return strings.Join([]string{"EnterpriseProject", string(data)}, " ")
}
