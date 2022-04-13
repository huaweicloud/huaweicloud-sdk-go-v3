package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type MfaTotp struct {
	User *MfaTotpUser `json:"user"`
}

func (o MfaTotp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaTotp struct{}"
	}

	return strings.Join([]string{"MfaTotp", string(data)}, " ")
}
