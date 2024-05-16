package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGroupResponse Response Object
type AddGroupResponse struct {

	// 群组ID
	Id *string `json:"id,omitempty"`

	// 群组名称
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGroupResponse struct{}"
	}

	return strings.Join([]string{"AddGroupResponse", string(data)}, " ")
}
