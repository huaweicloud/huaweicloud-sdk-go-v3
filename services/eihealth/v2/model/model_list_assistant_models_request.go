package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssistantModelsRequest Request Object
type ListAssistantModelsRequest struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	VendorId string `json:"vendor_id"`

	// **参数解释**： 排序规则，目前默认创建时间降序。 **约束限制**： 不涉及 **取值范围**：   * service_name：服务名称   * name：模型名称   * type：模型类型   * update_time：修改时间 **默认取值**： update_time
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序方向。 **约束限制**： 不涉及 **取值范围**： - acs：升序 - desc：降序 **默认取值**： desc
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAssistantModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssistantModelsRequest struct{}"
	}

	return strings.Join([]string{"ListAssistantModelsRequest", string(data)}, " ")
}
