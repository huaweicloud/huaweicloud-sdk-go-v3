package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardRequest Request Object
type UpdateStandardRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	Body *StandElementValueVoList `json:"body,omitempty"`
}

func (o UpdateStandardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardRequest struct{}"
	}

	return strings.Join([]string{"UpdateStandardRequest", string(data)}, " ")
}
