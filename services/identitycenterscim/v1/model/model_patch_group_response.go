package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchGroupResponse Response Object
type PatchGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PatchGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchGroupResponse struct{}"
	}

	return strings.Join([]string{"PatchGroupResponse", string(data)}, " ")
}
