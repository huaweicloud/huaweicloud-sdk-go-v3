package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExpansionParametersResponse Response Object
type ShowExpansionParametersResponse struct {

	// **参数解释**: 参数信息。
	ExpansionParameters *[]ExpansionParameterResult `json:"expansion_parameters,omitempty"`
	HttpStatusCode      int                         `json:"-"`
}

func (o ShowExpansionParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpansionParametersResponse struct{}"
	}

	return strings.Join([]string{"ShowExpansionParametersResponse", string(data)}, " ")
}
