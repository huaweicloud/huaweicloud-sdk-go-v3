package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationUpdateRequest struct {

	// 修改的应用名称。
	Name string `json:"name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 是否收藏。
	IsCollection *bool `json:"is_collection,omitempty"`
}

func (o ApplicationUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationUpdateRequest struct{}"
	}

	return strings.Join([]string{"ApplicationUpdateRequest", string(data)}, " ")
}
