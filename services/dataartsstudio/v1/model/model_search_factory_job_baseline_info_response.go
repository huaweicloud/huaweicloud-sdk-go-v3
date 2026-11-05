package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchFactoryJobBaselineInfoResponse Response Object
type SearchFactoryJobBaselineInfoResponse struct {

	// **参数解释**： 作业基线归属信息列表。 **取值范围**： 不涉及。
	Body           *[]JobBaselineInfo `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o SearchFactoryJobBaselineInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFactoryJobBaselineInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchFactoryJobBaselineInfoResponse", string(data)}, " ")
}
