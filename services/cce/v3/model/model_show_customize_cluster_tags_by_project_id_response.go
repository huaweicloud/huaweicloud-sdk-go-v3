package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomizeClusterTagsByProjectIdResponse Response Object
type ShowCustomizeClusterTagsByProjectIdResponse struct {

	// 资源标签
	Tags *[]ResourceTagBody `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]ResourceTagBody `json:"sys_tags,omitempty"`

	// 执行动作
	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomizeClusterTagsByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomizeClusterTagsByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomizeClusterTagsByProjectIdResponse", string(data)}, " ")
}
