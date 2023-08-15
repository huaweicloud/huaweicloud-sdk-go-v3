package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceByTagResponse Response Object
type ListResourceInstanceByTagResponse struct {

	// 资源实例列表
	Resources *[]ResourceInstanceResponseResources `json:"resources,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceInstanceByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceByTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceByTagResponse", string(data)}, " ")
}
