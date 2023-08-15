package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezeSubCustomersRequest Request Object
type FreezeSubCustomersRequest struct {
	Body *FreezeSubCustomersReq `json:"body,omitempty"`
}

func (o FreezeSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"FreezeSubCustomersRequest", string(data)}, " ")
}
