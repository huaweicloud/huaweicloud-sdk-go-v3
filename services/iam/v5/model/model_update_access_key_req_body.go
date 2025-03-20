package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessKeyReqBody The request body of update access key requests.
type UpdateAccessKeyReqBody struct {
	Status *AccessKeyStatus `json:"status"`
}

func (o UpdateAccessKeyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessKeyReqBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessKeyReqBody", string(data)}, " ")
}
