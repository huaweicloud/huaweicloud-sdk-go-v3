package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecialConfigurationRequest Request Object
type ListSpecialConfigurationRequest struct {

	// **参数解释：** 加速域名  **约束限制：** 仅支持查询已经在CDN添加成功的域名 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 每页加速域名的数量 **约束限制：** 不涉及 **取值范围：** 1-10000 **默认取值：** 30
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释：** 查询的页码 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 1
	PageNumber *int32 `json:"page_number,omitempty"`
}

func (o ListSpecialConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecialConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ListSpecialConfigurationRequest", string(data)}, " ")
}
