package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCceIntegrationProtectionResponse Response Object
type AddCceIntegrationProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddCceIntegrationProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCceIntegrationProtectionResponse struct{}"
	}

	return strings.Join([]string{"AddCceIntegrationProtectionResponse", string(data)}, " ")
}
