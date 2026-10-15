package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFlavorReq 规格变更请求体
type ChangeFlavorReq struct {

	// 目标规格类型。 取值： - micro：极小规格。 - small：小规格。 - medium：中规格。 - high：大规格。 - xlarge：超大规格。 - 2xlarge：极大规格。
	SpecType string `json:"spec_type"`
}

func (o ChangeFlavorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFlavorReq struct{}"
	}

	return strings.Join([]string{"ChangeFlavorReq", string(data)}, " ")
}
