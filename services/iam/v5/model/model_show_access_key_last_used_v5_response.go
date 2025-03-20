package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessKeyLastUsedV5Response Response Object
type ShowAccessKeyLastUsedV5Response struct {
	AccessKeyLastUsed *AccessKeyLastUsed `json:"access_key_last_used,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowAccessKeyLastUsedV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessKeyLastUsedV5Response struct{}"
	}

	return strings.Join([]string{"ShowAccessKeyLastUsedV5Response", string(data)}, " ")
}
