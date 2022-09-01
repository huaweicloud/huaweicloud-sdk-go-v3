package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteIpListRequest struct {

	// IP地址组ID。
	IpgroupId string `json:"ipgroup_id" xml:"ipgroup_id"`

	Body *BatchDeleteIpListRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListRequest", string(data)}, " ")
}
