package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationTemplatesRequest Request Object
type ListApplicationTemplatesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 应用程序ID，以app-为前缀
	ApplicationId string `json:"application_id"`
}

func (o ListApplicationTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationTemplatesRequest", string(data)}, " ")
}
