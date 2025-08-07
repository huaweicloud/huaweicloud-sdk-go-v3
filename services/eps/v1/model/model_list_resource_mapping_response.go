package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceMappingResponse Response Object
type ListResourceMappingResponse struct {

	// 资源类型映射
	ResourceMapping map[string]string `json:"resource_mapping,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListResourceMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceMappingResponse struct{}"
	}

	return strings.Join([]string{"ListResourceMappingResponse", string(data)}, " ")
}
