package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableModelRequest Request Object
type CreateTableModelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *TableModelVo `json:"body,omitempty"`
}

func (o CreateTableModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableModelRequest struct{}"
	}

	return strings.Join([]string{"CreateTableModelRequest", string(data)}, " ")
}
