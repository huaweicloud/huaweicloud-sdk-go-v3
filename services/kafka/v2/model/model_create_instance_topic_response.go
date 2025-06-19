package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceTopicResponse Response Object
type CreateInstanceTopicResponse struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicResponse", string(data)}, " ")
}
