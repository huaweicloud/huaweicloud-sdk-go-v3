package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataServiceInstanceLtsLogResponse Response Object
type UpdateDataServiceInstanceLtsLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDataServiceInstanceLtsLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataServiceInstanceLtsLogResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataServiceInstanceLtsLogResponse", string(data)}, " ")
}
