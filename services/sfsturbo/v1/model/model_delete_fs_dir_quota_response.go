package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsDirQuotaResponse Response Object
type DeleteFsDirQuotaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFsDirQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirQuotaResponse struct{}"
	}

	return strings.Join([]string{"DeleteFsDirQuotaResponse", string(data)}, " ")
}
