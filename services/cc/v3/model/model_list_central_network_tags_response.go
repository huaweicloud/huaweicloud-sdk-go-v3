package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkTagsResponse Response Object
type ListCentralNetworkTagsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 所有标签。
	Tags           []MultivaluedTag `json:"tags"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCentralNetworkTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkTagsResponse", string(data)}, " ")
}
