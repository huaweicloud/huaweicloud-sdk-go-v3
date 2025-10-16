package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPublicationRequest Request Object
type ModifyPublicationRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 发布ID
	PublicationId string `json:"publication_id"`

	Body *ModifyPublicationsRequestBody `json:"body,omitempty"`
}

func (o ModifyPublicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicationRequest struct{}"
	}

	return strings.Join([]string{"ModifyPublicationRequest", string(data)}, " ")
}
