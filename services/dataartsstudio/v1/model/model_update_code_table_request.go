package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableRequest Request Object
type UpdateCodeTableRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	Body *CodeTableVo `json:"body,omitempty"`
}

func (o UpdateCodeTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableRequest", string(data)}, " ")
}
