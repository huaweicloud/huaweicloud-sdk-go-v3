package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFileResponse struct {

	// 文件路径。
	Path           *string `json:"path,omitempty" xml:"path"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFileResponse struct{}"
	}

	return strings.Join([]string{"UpdateFileResponse", string(data)}, " ")
}
