package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFindingsResponse Response Object
type UpdateFindingsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFindingsResponse struct{}"
	}

	return strings.Join([]string{"UpdateFindingsResponse", string(data)}, " ")
}
