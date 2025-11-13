package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyResponse Response Object
type CreateAgencyResponse struct {

	// 创建OBS委托，由IAM返回的授权ID。
	AgencyId       *string `json:"agency_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyResponse", string(data)}, " ")
}
