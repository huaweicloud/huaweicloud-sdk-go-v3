package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationInstancesResponse Response Object
type ListApplicationInstancesResponse struct {

	// 应用程序实例列表
	ApplicationInstances *[]ApplicationInstanceDto `json:"application_instances,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApplicationInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationInstancesResponse", string(data)}, " ")
}
