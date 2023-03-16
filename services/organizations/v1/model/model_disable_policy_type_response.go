package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisablePolicyTypeResponse struct {
	Root           *RootDto `json:"root,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o DisablePolicyTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisablePolicyTypeResponse struct{}"
	}

	return strings.Join([]string{"DisablePolicyTypeResponse", string(data)}, " ")
}
