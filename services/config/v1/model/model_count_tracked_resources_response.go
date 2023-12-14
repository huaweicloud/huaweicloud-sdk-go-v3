package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTrackedResourcesResponse Response Object
type CountTrackedResourcesResponse struct {

	// 资源总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountTrackedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTrackedResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountTrackedResourcesResponse", string(data)}, " ")
}
