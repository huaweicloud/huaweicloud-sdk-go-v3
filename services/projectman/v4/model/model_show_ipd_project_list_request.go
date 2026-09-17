package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdProjectListRequest Request Object
type ShowIpdProjectListRequest struct {

	// **参数解释**： 项目名称搜索关键字。 **约束限制**： 最大256个字符。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Search *string `json:"search,omitempty"`

	// **参数解释**： IPD项目模型id。 **约束限制**： 不涉及 **取值范围**： 10001（系统设备类） 10002（独立软件类） 10003（云服务类型） **默认取值**： 不涉及
	Model *string `json:"model,omitempty"`
}

func (o ShowIpdProjectListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProjectListRequest struct{}"
	}

	return strings.Join([]string{"ShowIpdProjectListRequest", string(data)}, " ")
}
