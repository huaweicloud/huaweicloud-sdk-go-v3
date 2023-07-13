package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateApiRequest Request Object
type MigrateApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ApiMoveParaDto `json:"body,omitempty"`
}

func (o MigrateApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateApiRequest struct{}"
	}

	return strings.Join([]string{"MigrateApiRequest", string(data)}, " ")
}
