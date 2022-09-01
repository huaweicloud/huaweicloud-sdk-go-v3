package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOrgInstancesResponse struct {
	Instances *PageInstancesVo `json:"instances,omitempty" xml:"instances"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOrgInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrgInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListOrgInstancesResponse", string(data)}, " ")
}
