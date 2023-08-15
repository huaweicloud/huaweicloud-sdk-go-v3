package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPythonFilesRequest Request Object
type UploadPythonFilesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadGroupPackageReq `json:"body,omitempty"`
}

func (o UploadPythonFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPythonFilesRequest struct{}"
	}

	return strings.Join([]string{"UploadPythonFilesRequest", string(data)}, " ")
}
