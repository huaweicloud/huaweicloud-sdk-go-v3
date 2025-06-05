package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopProjectResponse Response Object
type UpdateTopProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopProjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopProjectResponse", string(data)}, " ")
}
