package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgenciesPrimitiveTypeHolder struct {

	// 委托授权的信息。
	Agencies *[]Agency `json:"agencies,omitempty"`
}

func (o AgenciesPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgenciesPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"AgenciesPrimitiveTypeHolder", string(data)}, " ")
}
