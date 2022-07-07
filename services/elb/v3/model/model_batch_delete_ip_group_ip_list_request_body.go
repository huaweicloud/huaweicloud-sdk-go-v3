package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Response Object
type BatchDeleteIpGroupIpListRequestBody struct {
	Ipgroup *BatchDeleteIpListOption `json:"ipgroup,omitempty"`
}

func (o BatchDeleteIpGroupIpListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpGroupIpListRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpGroupIpListRequestBody", string(data)}, " ")
}
