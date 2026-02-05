package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupSelectionResponse Response Object
type DeleteBackupSelectionResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBackupSelectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupSelectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupSelectionResponse", string(data)}, " ")
}
