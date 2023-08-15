package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceUserOrGroupResponse Response Object
type UpdateWorkSpaceUserOrGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWorkSpaceUserOrGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceUserOrGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceUserOrGroupResponse", string(data)}, " ")
}
