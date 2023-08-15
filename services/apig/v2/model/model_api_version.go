package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiVersion struct {

	// API版本的编号
	VersionId *string `json:"version_id,omitempty"`
}

func (o ApiVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersion struct{}"
	}

	return strings.Join([]string{"ApiVersion", string(data)}, " ")
}
