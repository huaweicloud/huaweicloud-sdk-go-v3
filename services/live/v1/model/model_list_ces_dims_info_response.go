package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesDimsInfoResponse Response Object
type ListCesDimsInfoResponse struct {
	Query          *CesQueryRespQuery `json:"query,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCesDimsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesDimsInfoResponse struct{}"
	}

	return strings.Join([]string{"ListCesDimsInfoResponse", string(data)}, " ")
}
