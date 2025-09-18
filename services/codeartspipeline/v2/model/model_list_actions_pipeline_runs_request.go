package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsPipelineRunsRequest Request Object
type ListActionsPipelineRunsRequest struct {

	// **参数解释**： 租户id。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	DomainId string `json:"domain_id"`

	Body *ActionsPipelineRunsQueryDto `json:"body,omitempty"`
}

func (o ListActionsPipelineRunsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsPipelineRunsRequest struct{}"
	}

	return strings.Join([]string{"ListActionsPipelineRunsRequest", string(data)}, " ")
}
