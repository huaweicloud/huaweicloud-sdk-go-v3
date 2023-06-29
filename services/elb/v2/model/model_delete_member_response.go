package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberResponse Response Object
type DeleteMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberResponse", string(data)}, " ")
}
