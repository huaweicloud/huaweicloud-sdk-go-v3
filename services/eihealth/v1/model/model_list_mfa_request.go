package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMfaRequest struct {

	// 用户id
	UserId string `json:"user_id"`
}

func (o ListMfaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMfaRequest struct{}"
	}

	return strings.Join([]string{"ListMfaRequest", string(data)}, " ")
}
