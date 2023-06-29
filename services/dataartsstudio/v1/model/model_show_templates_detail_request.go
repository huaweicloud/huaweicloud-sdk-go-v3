package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplatesDetailRequest Request Object
type ShowTemplatesDetailRequest struct {

	// id
	Id int64 `json:"id"`

	// workspace 信息
	Workspace string `json:"workspace"`
}

func (o ShowTemplatesDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplatesDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplatesDetailRequest", string(data)}, " ")
}
