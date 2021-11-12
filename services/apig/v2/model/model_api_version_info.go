package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiVersionInfo struct {
	// API版本的编号

	VersionId *string `json:"version_id,omitempty"`
}

func (o ApiVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionInfo struct{}"
	}

	return strings.Join([]string{"ApiVersionInfo", string(data)}, " ")
}
