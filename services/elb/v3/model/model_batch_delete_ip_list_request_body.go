package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Response Object
type BatchDeleteIpListRequestBody struct {
	Ipgroup *BatchDeleteIpListOption `json:"ipgroup,omitempty" xml:"ipgroup"`
}

func (o BatchDeleteIpListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListRequestBody", string(data)}, " ")
}
