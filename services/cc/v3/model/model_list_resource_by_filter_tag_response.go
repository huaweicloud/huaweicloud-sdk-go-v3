package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceByFilterTagResponse struct {

	// 资源实例实例列表。
	Resources *[]FilterTagResource `json:"resources,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// 符合过滤条件的资源总数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceByFilterTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByFilterTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourceByFilterTagResponse", string(data)}, " ")
}
