package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDedicatedHostRequest struct {
	Body *ReqAllocateDeh `json:"body,omitempty"`
}

func (o CreateDedicatedHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDedicatedHostRequest struct{}"
	}

	return strings.Join([]string{"CreateDedicatedHostRequest", string(data)}, " ")
}
