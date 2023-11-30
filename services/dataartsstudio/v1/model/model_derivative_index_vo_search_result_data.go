package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DerivativeIndexVoSearchResultData struct {
	Value *DerivativeIndexVoSearchResultDataValue `json:"value,omitempty"`
}

func (o DerivativeIndexVoSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexVoSearchResultData struct{}"
	}

	return strings.Join([]string{"DerivativeIndexVoSearchResultData", string(data)}, " ")
}
