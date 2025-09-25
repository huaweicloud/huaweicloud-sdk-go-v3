package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedistributionParametersResponse Response Object
type ShowRedistributionParametersResponse struct {

	// **参数解释**: 参数列表信息。
	RedistributionParameters *[]RedistributionParameterResult `json:"redistribution_parameters,omitempty"`
	HttpStatusCode           int                              `json:"-"`
}

func (o ShowRedistributionParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedistributionParametersResponse struct{}"
	}

	return strings.Join([]string{"ShowRedistributionParametersResponse", string(data)}, " ")
}
