package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmResourcesResponse struct {
	// 资源信息

	Resources *[][]Dimension `json:"resources,omitempty"`
	// 资源总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmResourcesResponse", string(data)}, " ")
}
