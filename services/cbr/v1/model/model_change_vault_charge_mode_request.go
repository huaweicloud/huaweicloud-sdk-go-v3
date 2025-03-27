package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVaultChargeModeRequest Request Object
type ChangeVaultChargeModeRequest struct {
	Body *ChangeToPeriod `json:"body,omitempty"`
}

func (o ChangeVaultChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVaultChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeVaultChargeModeRequest", string(data)}, " ")
}
