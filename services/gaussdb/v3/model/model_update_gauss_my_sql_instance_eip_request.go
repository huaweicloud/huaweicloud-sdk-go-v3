package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceEipRequest Request Object
type UpdateGaussMySqlInstanceEipRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyBindEipRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceEipRequest", string(data)}, " ")
}
