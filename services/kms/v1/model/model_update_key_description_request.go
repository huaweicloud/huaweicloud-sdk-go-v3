package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateKeyDescriptionRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *UpdateKeyDescriptionRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateKeyDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyDescriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyDescriptionRequest", string(data)}, " ")
}
