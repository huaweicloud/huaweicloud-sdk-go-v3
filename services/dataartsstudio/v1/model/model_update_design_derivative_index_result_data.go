package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignDerivativeIndexResultData 更新衍生指标的返回结果，成功的个数。
type UpdateDesignDerivativeIndexResultData struct {
	Value *DerivativeIndexVo `json:"value,omitempty"`
}

func (o UpdateDesignDerivativeIndexResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignDerivativeIndexResultData struct{}"
	}

	return strings.Join([]string{"UpdateDesignDerivativeIndexResultData", string(data)}, " ")
}
