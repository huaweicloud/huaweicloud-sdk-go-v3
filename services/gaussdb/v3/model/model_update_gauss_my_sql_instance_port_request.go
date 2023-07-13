package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstancePortRequest Request Object
type UpdateGaussMySqlInstancePortRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyPortRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstancePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstancePortRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstancePortRequest", string(data)}, " ")
}
