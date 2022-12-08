package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CountResourceInstanceByTagResponse struct {

	// 资源实例列表
	Resources *[]ResourceInstanceResponseResources `json:"resources,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountResourceInstanceByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstanceByTagResponse struct{}"
	}

	return strings.Join([]string{"CountResourceInstanceByTagResponse", string(data)}, " ")
}
