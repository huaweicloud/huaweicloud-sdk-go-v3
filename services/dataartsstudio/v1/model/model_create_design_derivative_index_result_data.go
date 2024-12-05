package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignDerivativeIndexResultData 创建衍生指标的返回结果，成功的个数。
type CreateDesignDerivativeIndexResultData struct {
	Value *BatchOperationVo `json:"value,omitempty"`
}

func (o CreateDesignDerivativeIndexResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignDerivativeIndexResultData struct{}"
	}

	return strings.Join([]string{"CreateDesignDerivativeIndexResultData", string(data)}, " ")
}
