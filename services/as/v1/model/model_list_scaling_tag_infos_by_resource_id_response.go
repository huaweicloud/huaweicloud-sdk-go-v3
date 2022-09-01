package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScalingTagInfosByResourceIdResponse struct {

	// 资源标签列表。
	Tags *[]TagsSingleValue `json:"tags,omitempty" xml:"tags"`

	// 系统资源标签列表。
	SysTags        *[]TagsSingleValue `json:"sys_tags,omitempty" xml:"sys_tags"`
	HttpStatusCode int                `json:"-"`
}

func (o ListScalingTagInfosByResourceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByResourceIdResponse struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByResourceIdResponse", string(data)}, " ")
}
