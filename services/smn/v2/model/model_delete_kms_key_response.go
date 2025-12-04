package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKmsKeyResponse Response Object
type DeleteKmsKeyResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKmsKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKmsKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteKmsKeyResponse", string(data)}, " ")
}
