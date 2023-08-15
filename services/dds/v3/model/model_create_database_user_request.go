package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseUserRequest Request Object
type CreateDatabaseUserRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseUserRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequest", string(data)}, " ")
}
