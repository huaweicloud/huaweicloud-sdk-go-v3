package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAutoOpenQuotaStatusResponse Response Object
type ChangeAutoOpenQuotaStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAutoOpenQuotaStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAutoOpenQuotaStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAutoOpenQuotaStatusResponse", string(data)}, " ")
}
