package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveResourceLevelResponse Response Object
type RemoveResourceLevelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveResourceLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveResourceLevelResponse struct{}"
	}

	return strings.Join([]string{"RemoveResourceLevelResponse", string(data)}, " ")
}
