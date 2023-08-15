package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableValuesRequest Request Object
type UpdateCodeTableValuesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	Body *CodeTableFieldValueUpdateVo `json:"body,omitempty"`
}

func (o UpdateCodeTableValuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableValuesRequest struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableValuesRequest", string(data)}, " ")
}
