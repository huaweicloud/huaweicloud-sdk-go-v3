package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableCcspInstanceResponse Response Object
type DisableCcspInstanceResponse struct {

	// 实例ID
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableCcspInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCcspInstanceResponse struct{}"
	}

	return strings.Join([]string{"DisableCcspInstanceResponse", string(data)}, " ")
}
