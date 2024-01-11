package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBackup2Response Response Object
type ExportBackup2Response struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportBackup2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBackup2Response struct{}"
	}

	return strings.Join([]string{"ExportBackup2Response", string(data)}, " ")
}
