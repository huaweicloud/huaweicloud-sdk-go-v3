package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVulnerabilitiesResponse Response Object
type ExportVulnerabilitiesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ExportVulnerabilitiesResponse", string(data)}, " ")
}
