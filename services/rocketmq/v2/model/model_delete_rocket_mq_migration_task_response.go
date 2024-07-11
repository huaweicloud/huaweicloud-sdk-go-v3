package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRocketMqMigrationTaskResponse Response Object
type DeleteRocketMqMigrationTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteRocketMqMigrationTaskResponse", string(data)}, " ")
}
