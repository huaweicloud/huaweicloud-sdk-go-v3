package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGaussMySqlInstanceOpsWindowRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID
	InstanceId string `json:"instance_id"`

	Body *ModifyOpsWindow `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceOpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceOpsWindowRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceOpsWindowRequest", string(data)}, " ")
}
