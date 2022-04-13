package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateParametersForImportRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *GetParametersForImportRequestBody `json:"body,omitempty"`
}

func (o CreateParametersForImportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParametersForImportRequest struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportRequest", string(data)}, " ")
}
