package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageEmailConfigResponse Response Object
type UpdateMessageEmailConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMessageEmailConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageEmailConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateMessageEmailConfigResponse", string(data)}, " ")
}
