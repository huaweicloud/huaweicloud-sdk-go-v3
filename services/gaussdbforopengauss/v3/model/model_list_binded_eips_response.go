package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBindedEipsResponse Response Object
type ListBindedEipsResponse struct {

	// 查询实例已绑定EIP列表。
	PublicIps *[]BindedEipResult `json:"public_ips,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBindedEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindedEipsResponse struct{}"
	}

	return strings.Join([]string{"ListBindedEipsResponse", string(data)}, " ")
}
