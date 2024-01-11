package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBackup2Response Response Object
type ImportBackup2Response struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportBackup2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackup2Response struct{}"
	}

	return strings.Join([]string{"ImportBackup2Response", string(data)}, " ")
}
