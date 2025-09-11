package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResourceInstancesCountResponse Response Object
type ShowInstanceResourceInstancesCountResponse struct {

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceResourceInstancesCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResourceInstancesCountResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResourceInstancesCountResponse", string(data)}, " ")
}
