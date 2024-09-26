package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionTagsResponse Response Object
type ListCloudConnectionTagsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 云连接实例的所有标签。
	Tags           []MultivaluedTag `json:"tags"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCloudConnectionTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionTagsResponse", string(data)}, " ")
}
