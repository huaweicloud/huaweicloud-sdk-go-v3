package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateIpListRequest struct {

	// IP地址组ID。
	IpgroupId string `json:"ipgroup_id" xml:"ipgroup_id"`

	Body *UpdateIpListRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequest", string(data)}, " ")
}
