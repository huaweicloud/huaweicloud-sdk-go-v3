package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpListRequestBody This is a auto create Body Object
type UpdateIpListRequestBody struct {
	Ipgroup *UpdateIpListOption `json:"ipgroup,omitempty"`
}

func (o UpdateIpListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequestBody", string(data)}, " ")
}
