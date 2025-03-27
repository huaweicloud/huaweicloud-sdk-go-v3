package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigprojectResponse Response Object
type DeleteMigprojectResponse struct {

	// 删除迁移项目成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigprojectResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectResponse", string(data)}, " ")
}
