package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyForAccessRequest Request Object
type ApplyForAccessRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`

	// 申请接入服务的请求信息
	Body *[]AccessRequestInfo `json:"body,omitempty"`
}

func (o ApplyForAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyForAccessRequest struct{}"
	}

	return strings.Join([]string{"ApplyForAccessRequest", string(data)}, " ")
}
