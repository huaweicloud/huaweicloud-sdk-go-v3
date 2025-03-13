package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSecurityDiagnoseResponse Response Object
type ExecuteSecurityDiagnoseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteSecurityDiagnoseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSecurityDiagnoseResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSecurityDiagnoseResponse", string(data)}, " ")
}
