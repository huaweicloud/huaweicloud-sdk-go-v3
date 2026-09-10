package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceOldResponse Response Object
type UpdateWorkSpaceOldResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWorkSpaceOldResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceOldResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceOldResponse", string(data)}, " ")
}
