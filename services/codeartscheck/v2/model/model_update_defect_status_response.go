package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefectStatusResponse Response Object
type UpdateDefectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDefectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefectStatusResponse", string(data)}, " ")
}
