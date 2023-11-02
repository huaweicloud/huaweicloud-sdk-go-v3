package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableRestoreResponse Response Object
type CheckTableRestoreResponse struct {
	CheckTableNameResult *CheckTableNameResult `json:"check_table_name_result,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o CheckTableRestoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableRestoreResponse struct{}"
	}

	return strings.Join([]string{"CheckTableRestoreResponse", string(data)}, " ")
}
