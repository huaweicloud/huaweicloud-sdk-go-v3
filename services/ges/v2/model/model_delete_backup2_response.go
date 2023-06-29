package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackup2Response Response Object
type DeleteBackup2Response struct {
	Body           map[string]string `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteBackup2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackup2Response struct{}"
	}

	return strings.Join([]string{"DeleteBackup2Response", string(data)}, " ")
}
