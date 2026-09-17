package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionDomainInfoDetail 特性白名单响应体。
type ActionDomainInfoDetail struct {

	// **参数解释**： 特性ID。 **取值范围**： 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**： 特性描述。当请求头x-language为en-us时为英文，为zh-cn或不传则返回描述信息为中文。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ActionDomainInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionDomainInfoDetail struct{}"
	}

	return strings.Join([]string{"ActionDomainInfoDetail", string(data)}, " ")
}
