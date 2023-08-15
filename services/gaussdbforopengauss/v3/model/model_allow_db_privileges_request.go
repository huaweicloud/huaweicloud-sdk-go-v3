package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowDbPrivilegesRequest Request Object
type AllowDbPrivilegesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *GaussDBforOpenGaussGrantRequest `json:"body,omitempty"`
}

func (o AllowDbPrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbPrivilegesRequest struct{}"
	}

	return strings.Join([]string{"AllowDbPrivilegesRequest", string(data)}, " ")
}
