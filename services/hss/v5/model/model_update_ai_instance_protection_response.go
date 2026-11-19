package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiInstanceProtectionResponse Response Object
type UpdateAiInstanceProtectionResponse struct {

	// **参数解释**： 每个智能体的处理结果列表 **取值范围**： 不涉及
	Results        *[]UpdateAiInstanceProtectionResultItem `json:"results,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o UpdateAiInstanceProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiInstanceProtectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiInstanceProtectionResponse", string(data)}, " ")
}
