package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEipStatusResponse Response Object
type ChangeEipStatusResponse struct {
	Data           *EipSwitchStatusVo `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ChangeEipStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEipStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeEipStatusResponse", string(data)}, " ")
}
