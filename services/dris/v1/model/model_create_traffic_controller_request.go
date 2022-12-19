package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTrafficControllerRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddTrafficControllerDto `json:"body,omitempty"`
}

func (o CreateTrafficControllerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrafficControllerRequest struct{}"
	}

	return strings.Join([]string{"CreateTrafficControllerRequest", string(data)}, " ")
}
