package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefaultMigprojectResponse Response Object
type UpdateDefaultMigprojectResponse struct {

	// 更改默认迁移项目成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDefaultMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectResponse", string(data)}, " ")
}
