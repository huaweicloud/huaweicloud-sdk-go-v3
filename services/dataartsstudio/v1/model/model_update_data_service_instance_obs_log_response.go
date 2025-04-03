package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataServiceInstanceObsLogResponse Response Object
type UpdateDataServiceInstanceObsLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDataServiceInstanceObsLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataServiceInstanceObsLogResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataServiceInstanceObsLogResponse", string(data)}, " ")
}
