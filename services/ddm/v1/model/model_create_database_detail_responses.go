package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseDetailResponses This is a auto create Response Object
type CreateDatabaseDetailResponses struct {

	// 逻辑库名称。
	Name string `json:"name"`
}

func (o CreateDatabaseDetailResponses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseDetailResponses struct{}"
	}

	return strings.Join([]string{"CreateDatabaseDetailResponses", string(data)}, " ")
}
