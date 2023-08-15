package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupRequestBody This is a auto create Body Object
type UpdateIpGroupRequestBody struct {
	Ipgroup *UpdateIpGroupOption `json:"ipgroup"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
