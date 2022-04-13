package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportKeyMaterialRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *ImportKeyMaterialRequestBody `json:"body,omitempty"`
}

func (o ImportKeyMaterialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportKeyMaterialRequest struct{}"
	}

	return strings.Join([]string{"ImportKeyMaterialRequest", string(data)}, " ")
}
