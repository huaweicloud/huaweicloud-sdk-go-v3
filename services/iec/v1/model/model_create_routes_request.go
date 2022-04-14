package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRoutesRequest struct {
	// 路由表ID

	RoutetableId string `json:"routetable_id"`

	Body *CreateRoutesRequestBody `json:"body,omitempty"`
}

func (o CreateRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutesRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutesRequest", string(data)}, " ")
}
