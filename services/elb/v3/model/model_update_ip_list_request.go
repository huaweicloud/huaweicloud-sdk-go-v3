package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpListRequest Request Object
type UpdateIpListRequest struct {

	// 参数解释：IP地址组ID。
	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpListRequestBody `json:"body,omitempty"`
}

func (o UpdateIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequest", string(data)}, " ")
}
