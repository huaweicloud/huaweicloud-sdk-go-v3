package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretsConfigRequest Request Object
type ShowSecretsConfigRequest struct {
}

func (o ShowSecretsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretsConfigRequest", string(data)}, " ")
}
