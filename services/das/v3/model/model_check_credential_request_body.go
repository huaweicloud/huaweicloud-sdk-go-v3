package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckCredentialRequestBody struct {

	// OBS桶名
	BucketName string `json:"bucket_name"`

	// AK
	Ak string `json:"ak"`

	// SK
	Sk string `json:"sk"`
}

func (o CheckCredentialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCredentialRequestBody struct{}"
	}

	return strings.Join([]string{"CheckCredentialRequestBody", string(data)}, " ")
}
