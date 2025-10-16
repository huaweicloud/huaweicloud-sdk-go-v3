package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicationRequest Request Object
type DeletePublicationRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeletePublicationRequestBody `json:"body,omitempty"`
}

func (o DeletePublicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicationRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicationRequest", string(data)}, " ")
}
