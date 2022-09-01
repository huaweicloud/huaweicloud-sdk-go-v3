package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type FreezeSubCustomersRequest struct {
	Body *FreezeSubCustomersReq `json:"body,omitempty" xml:"body"`
}

func (o FreezeSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"FreezeSubCustomersRequest", string(data)}, " ")
}
