package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDatabaseUserRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseUserRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequest", string(data)}, " ")
}
