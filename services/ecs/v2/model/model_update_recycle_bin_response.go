package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinResponse Response Object
type UpdateRecycleBinResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRecycleBinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinResponse", string(data)}, " ")
}
