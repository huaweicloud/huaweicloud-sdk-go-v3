package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerChargeModeRequest Request Object
type ChangeServerChargeModeRequest struct {
	Body *ChangeServerChargeModeRequestBody `json:"body,omitempty"`
}

func (o ChangeServerChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerChargeModeRequest", string(data)}, " ")
}
