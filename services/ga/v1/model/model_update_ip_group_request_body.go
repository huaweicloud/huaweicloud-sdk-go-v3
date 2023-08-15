package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupRequestBody update ip group request
type UpdateIpGroupRequestBody struct {
	IpGroup *UpdateIpGroupOption `json:"ip_group"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
