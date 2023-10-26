package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateBackUpInfoRequest Request Object
type CreateOrUpdateBackUpInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AdditionalBackupRequest `json:"body,omitempty"`
}

func (o CreateOrUpdateBackUpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateBackUpInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateBackUpInfoRequest", string(data)}, " ")
}
