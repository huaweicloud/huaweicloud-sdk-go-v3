package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDnsMappingRequest Request Object
type UpdateUserDnsMappingRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *string `json:"body,omitempty"`
}

func (o UpdateUserDnsMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDnsMappingRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserDnsMappingRequest", string(data)}, " ")
}
