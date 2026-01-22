package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDomainResponseData struct {

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 项目ID **取值范围**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 域名信息列表 **取值范围**： 不涉及
	Records *[]DomainInfo `json:"records,omitempty"`

	// **参数解释**： 域名组id **取值范围**： 不涉及
	SetId *string `json:"set_id,omitempty"`

	// **参数解释**： 域名总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o ListDomainResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainResponseData struct{}"
	}

	return strings.Join([]string{"ListDomainResponseData", string(data)}, " ")
}
