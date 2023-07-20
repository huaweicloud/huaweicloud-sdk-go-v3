package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalValueReq
type UpdateGlobalValueReq struct {

	// 变量值
	VarValue string `json:"var_value"`
}

func (o UpdateGlobalValueReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalValueReq struct{}"
	}

	return strings.Join([]string{"UpdateGlobalValueReq", string(data)}, " ")
}
