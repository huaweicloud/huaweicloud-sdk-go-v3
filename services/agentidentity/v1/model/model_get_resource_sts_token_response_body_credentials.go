package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceStsTokenResponseBodyCredentials Credentials for API authentication
type GetResourceStsTokenResponseBodyCredentials struct {

	// The access key ID that identifies the temporary security credentials
	AccessKeyId string `json:"access_key_id"`

	// The date and time on which the current credentials expire
	Expiration *sdktime.SdkTime `json:"expiration"`

	// The secret access key that can be used to sign requests
	SecretAccessKey string `json:"secret_access_key"`

	// The token that users must pass to the service API to use the temporary credentials
	SecurityToken string `json:"security_token"`
}

func (o GetResourceStsTokenResponseBodyCredentials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceStsTokenResponseBodyCredentials struct{}"
	}

	return strings.Join([]string{"GetResourceStsTokenResponseBodyCredentials", string(data)}, " ")
}
