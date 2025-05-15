package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegionsResponse Response Object
type ListRegionsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 网络实例列表。
	Regions        []Region `json:"regions"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListRegionsResponse", string(data)}, " ")
}
