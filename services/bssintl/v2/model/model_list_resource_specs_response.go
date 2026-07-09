package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSpecsResponse Response Object
type ListResourceSpecsResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// |参数名称：资源规格信息列表| |参数的约束及描述：资源规格信息列表|
	CloudServiceBasics *[]CloudServiceBasic `json:"cloud_service_basics,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListResourceSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSpecsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceSpecsResponse", string(data)}, " ")
}
