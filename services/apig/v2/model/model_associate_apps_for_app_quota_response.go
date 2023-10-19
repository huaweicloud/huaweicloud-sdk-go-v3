package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAppsForAppQuotaResponse Response Object
type AssociateAppsForAppQuotaResponse struct {

	// 凭据与凭据配额绑定列表
	Applies        *[]AppQuotaAppBinding `json:"applies,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o AssociateAppsForAppQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsForAppQuotaResponse struct{}"
	}

	return strings.Join([]string{"AssociateAppsForAppQuotaResponse", string(data)}, " ")
}
