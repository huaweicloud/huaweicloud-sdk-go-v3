package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoutesRequest Request Object
type DeleteRoutesRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`

	Body *DeleteRoutesRequestBody `json:"body,omitempty"`
}

func (o DeleteRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutesRequest struct{}"
	}

	return strings.Join([]string{"DeleteRoutesRequest", string(data)}, " ")
}
