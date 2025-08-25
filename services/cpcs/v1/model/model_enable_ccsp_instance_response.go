package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCcspInstanceResponse Response Object
type EnableCcspInstanceResponse struct {

	// 实例ID
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableCcspInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCcspInstanceResponse struct{}"
	}

	return strings.Join([]string{"EnableCcspInstanceResponse", string(data)}, " ")
}
