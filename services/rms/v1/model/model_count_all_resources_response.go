package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllResourcesResponse Response Object
type CountAllResourcesResponse struct {

	// 资源总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountAllResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountAllResourcesResponse", string(data)}, " ")
}
