package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessKeyV5Response Response Object
type CreateAccessKeyV5Response struct {
	AccessKey      *AccessKey `json:"access_key,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateAccessKeyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessKeyV5Response struct{}"
	}

	return strings.Join([]string{"CreateAccessKeyV5Response", string(data)}, " ")
}
