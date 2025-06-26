package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubResourceInstancesCountResponse Response Object
type ShowSubResourceInstancesCountResponse struct {

	// 资源实例数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSubResourceInstancesCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubResourceInstancesCountResponse struct{}"
	}

	return strings.Join([]string{"ShowSubResourceInstancesCountResponse", string(data)}, " ")
}
