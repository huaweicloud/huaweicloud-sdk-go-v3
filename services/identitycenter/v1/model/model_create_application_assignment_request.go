package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationAssignmentRequest Request Object
type CreateApplicationAssignmentRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 应用程序实例ID，以app-ins-为前缀
	ApplicationInstanceId string `json:"application_instance_id"`

	Body *CreateApplicationAssignmentReqBody `json:"body,omitempty"`
}

func (o CreateApplicationAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationAssignmentRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationAssignmentRequest", string(data)}, " ")
}
