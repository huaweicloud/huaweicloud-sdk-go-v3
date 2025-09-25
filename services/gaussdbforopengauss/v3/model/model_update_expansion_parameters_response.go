package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExpansionParametersResponse Response Object
type UpdateExpansionParametersResponse struct {

	// **参数解释**: 修改扩容优化参数的任务ID。 **取值范围**: 不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateExpansionParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpansionParametersResponse struct{}"
	}

	return strings.Join([]string{"UpdateExpansionParametersResponse", string(data)}, " ")
}
