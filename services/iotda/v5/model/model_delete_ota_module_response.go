package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOtaModuleResponse Response Object
type DeleteOtaModuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOtaModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOtaModuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteOtaModuleResponse", string(data)}, " ")
}
