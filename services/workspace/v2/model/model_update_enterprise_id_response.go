package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseIdResponse Response Object
type UpdateEnterpriseIdResponse struct {

	// 企业ID。
	EnterpriseId   *string `json:"enterprise_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnterpriseIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseIdResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseIdResponse", string(data)}, " ")
}
