package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEchoTestPackageUsingResponse Response Object
type ShowEchoTestPackageUsingResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEchoTestPackageUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEchoTestPackageUsingResponse struct{}"
	}

	return strings.Join([]string{"ShowEchoTestPackageUsingResponse", string(data)}, " ")
}
