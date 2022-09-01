package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSecretForScheduleRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name" xml:"secret_name"`

	Body *DeleteSecretForScheduleRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteSecretForScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretForScheduleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretForScheduleRequest", string(data)}, " ")
}
