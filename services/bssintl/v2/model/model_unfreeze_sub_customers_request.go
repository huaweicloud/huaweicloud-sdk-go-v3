package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeSubCustomersRequest Request Object
type UnfreezeSubCustomersRequest struct {
	Body *UnfreezeSubCustomersReq `json:"body,omitempty"`
}

func (o UnfreezeSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"UnfreezeSubCustomersRequest", string(data)}, " ")
}
