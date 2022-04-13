package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddOrRemoveResourceInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrRemoveResourceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceResponse", string(data)}, " ")
}
