package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsTaskResponse Response Object
type DeleteFsTaskResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteFsTaskResponse", string(data)}, " ")
}
