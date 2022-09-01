package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AllowDbPrivilegesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *GaussDBforOpenGaussGrantRequest `json:"body,omitempty" xml:"body"`
}

func (o AllowDbPrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbPrivilegesRequest struct{}"
	}

	return strings.Join([]string{"AllowDbPrivilegesRequest", string(data)}, " ")
}
