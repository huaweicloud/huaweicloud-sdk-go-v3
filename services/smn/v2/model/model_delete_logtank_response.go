package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogtankResponse Response Object
type DeleteLogtankResponse struct {

	// 请求的唯一标识。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogtankResponse", string(data)}, " ")
}
