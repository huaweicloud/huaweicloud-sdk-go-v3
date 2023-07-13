package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoutesRequest Request Object
type UpdateRoutesRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`

	Body *UpdateRoutesRequestBody `json:"body,omitempty"`
}

func (o UpdateRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutesRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoutesRequest", string(data)}, " ")
}
