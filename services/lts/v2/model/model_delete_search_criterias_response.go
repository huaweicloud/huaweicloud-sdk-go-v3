package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSearchCriteriasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSearchCriteriasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSearchCriteriasResponse struct{}"
	}

	return strings.Join([]string{"DeleteSearchCriteriasResponse", string(data)}, " ")
}
