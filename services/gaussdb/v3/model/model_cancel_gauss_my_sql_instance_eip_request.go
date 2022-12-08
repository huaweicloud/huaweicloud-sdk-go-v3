package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelGaussMySqlInstanceEipRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例id
	InstanceId string `json:"instance_id"`
}

func (o CancelGaussMySqlInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGaussMySqlInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"CancelGaussMySqlInstanceEipRequest", string(data)}, " ")
}
