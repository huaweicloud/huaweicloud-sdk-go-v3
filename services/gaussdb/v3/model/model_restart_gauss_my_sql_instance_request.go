package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartGaussMySqlInstanceRequest Request Object
type RestartGaussMySqlInstanceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *TaurusRestartInstanceRequest `json:"body,omitempty"`
}

func (o RestartGaussMySqlInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGaussMySqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartGaussMySqlInstanceRequest", string(data)}, " ")
}
