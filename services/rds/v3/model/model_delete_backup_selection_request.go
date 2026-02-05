package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupSelectionRequest Request Object
type DeleteBackupSelectionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeleteBackupSelectionRequestBody `json:"body,omitempty"`
}

func (o DeleteBackupSelectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupSelectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupSelectionRequest", string(data)}, " ")
}
