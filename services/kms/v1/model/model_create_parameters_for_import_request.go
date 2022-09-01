package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateParametersForImportRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *GetParametersForImportRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateParametersForImportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParametersForImportRequest struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportRequest", string(data)}, " ")
}
