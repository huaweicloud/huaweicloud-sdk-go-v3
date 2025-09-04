package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableInnerSpecResponse Response Object
type ShowAvailableInnerSpecResponse struct {

	// **参数解释**： 内置执行机规格。 **取值范围**： 不涉及。
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAvailableInnerSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableInnerSpecResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableInnerSpecResponse", string(data)}, " ")
}
