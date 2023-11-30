package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DerivativeIndexVoDetailData struct {
	Value *DerivativeIndexVo `json:"value,omitempty"`
}

func (o DerivativeIndexVoDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexVoDetailData struct{}"
	}

	return strings.Join([]string{"DerivativeIndexVoDetailData", string(data)}, " ")
}
