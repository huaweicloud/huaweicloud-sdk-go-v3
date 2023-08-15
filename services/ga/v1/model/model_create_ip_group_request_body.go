package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupRequestBody create ip group request
type CreateIpGroupRequestBody struct {
	IpGroup *CreateIpGroupOption `json:"ip_group"`
}

func (o CreateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequestBody", string(data)}, " ")
}
