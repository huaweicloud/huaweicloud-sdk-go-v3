package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMultiResourcesResponse Response Object
type CountMultiResourcesResponse struct {

	// 云资源数量列表
	Data           *[]ResourceMultiCountResponseData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o CountMultiResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMultiResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountMultiResourcesResponse", string(data)}, " ")
}
