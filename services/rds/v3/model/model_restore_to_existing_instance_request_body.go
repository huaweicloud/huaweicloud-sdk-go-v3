package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreToExistingInstanceRequestBody struct {
	Source *RestoreToExistingInstanceRequestBodySource `json:"source" xml:"source"`

	Target *RestoreToExistingInstanceRequestBodyTarget `json:"target" xml:"target"`
}

func (o RestoreToExistingInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequestBody", string(data)}, " ")
}
