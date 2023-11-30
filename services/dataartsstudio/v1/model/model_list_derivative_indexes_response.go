package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDerivativeIndexesResponse Response Object
type ListDerivativeIndexesResponse struct {
	Data           *DerivativeIndexVoSearchResultData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListDerivativeIndexesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDerivativeIndexesResponse struct{}"
	}

	return strings.Join([]string{"ListDerivativeIndexesResponse", string(data)}, " ")
}
