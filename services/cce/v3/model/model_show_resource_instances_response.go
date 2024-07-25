package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceInstancesResponse Response Object
type ShowResourceInstancesResponse struct {

	// 数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源
	Resources      *[]ResInstanceBody `json:"resources,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceInstancesResponse", string(data)}, " ")
}
