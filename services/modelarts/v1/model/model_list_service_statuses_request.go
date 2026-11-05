package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceStatusesRequest Request Object
type ListServiceStatusesRequest struct {

	// **参数解释**：服务编码列表，逗号分隔，如aom,cce,ces,swr-ee。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ServiceCodes string `json:"service_codes"`
}

func (o ListServiceStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceStatusesRequest struct{}"
	}

	return strings.Join([]string{"ListServiceStatusesRequest", string(data)}, " ")
}
