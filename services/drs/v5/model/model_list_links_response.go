package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLinksResponse Response Object
type ListLinksResponse struct {

	// **参数解释**：  可用链路信息列表，包含每条可用链路的详细配置信息，详情请参见JobLinkResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	JobLinks *[]JobLinkResp `json:"job_links,omitempty"`

	// **参数解释**：  可用链路总条数，表示满足查询条件的链路记录总数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinksResponse struct{}"
	}

	return strings.Join([]string{"ListLinksResponse", string(data)}, " ")
}
