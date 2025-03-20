package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessKeyV5Response Response Object
type UpdateAccessKeyV5Response struct {
	AccessKey      *AccessKeyMetadata `json:"access_key,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateAccessKeyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessKeyV5Response struct{}"
	}

	return strings.Join([]string{"UpdateAccessKeyV5Response", string(data)}, " ")
}
