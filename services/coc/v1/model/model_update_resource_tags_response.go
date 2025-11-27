package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceTagsResponse Response Object
type UpdateResourceTagsResponse struct {

	// 更新脚本指定的资源标签，系统返回的脚本ID。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceTagsResponse", string(data)}, " ")
}
