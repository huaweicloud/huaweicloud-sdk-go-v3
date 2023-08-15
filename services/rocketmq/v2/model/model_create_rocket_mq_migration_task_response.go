package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRocketMqMigrationTaskResponse Response Object
type CreateRocketMqMigrationTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskResponse", string(data)}, " ")
}
