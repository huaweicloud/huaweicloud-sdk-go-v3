package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStorageRequest Request Object
type ShowStorageRequest struct {

	// 类型列表
	FormatList *string `json:"format_list,omitempty"`

	// 是否在项目中
	InProject *string `json:"in_project,omitempty"`
}

func (o ShowStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStorageRequest struct{}"
	}

	return strings.Join([]string{"ShowStorageRequest", string(data)}, " ")
}
