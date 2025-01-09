package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetIdsRequest Request Object
type UpdateSubnetIdsRequest struct {

	// 站点ID。
	SiteId string `json:"site_id"`

	Body *UpdateSubnetIdsRequestBody `json:"body,omitempty"`
}

func (o UpdateSubnetIdsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetIdsRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetIdsRequest", string(data)}, " ")
}
