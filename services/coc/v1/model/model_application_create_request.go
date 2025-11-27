package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationCreateRequest struct {

	// 应用名称。
	Name string `json:"name"`

	// 父节点id。
	ParentId *string `json:"parent_id,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`
}

func (o ApplicationCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationCreateRequest struct{}"
	}

	return strings.Join([]string{"ApplicationCreateRequest", string(data)}, " ")
}
