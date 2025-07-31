package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonTipsResponse Response Object
type ListCommonTipsResponse struct {
	HostName *CommonList `json:"host_name,omitempty"`

	HostId *CommonList `json:"host_id,omitempty"`

	PublicIp *CommonList `json:"public_ip,omitempty"`

	PrivateIp *CommonList `json:"private_ip,omitempty"`

	VpcId *CommonList `json:"vpc_id,omitempty"`

	ClusterId *CommonList `json:"cluster_id,omitempty"`

	HostTags       *HostTagInfoList `json:"host_tags,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCommonTipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonTipsResponse struct{}"
	}

	return strings.Join([]string{"ListCommonTipsResponse", string(data)}, " ")
}
