package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortalInfoResponseModel 更新主页响应体。
type UpdatePortalInfoResponseModel struct {
	Portal *PortalModel `json:"portal,omitempty"`
}

func (o UpdatePortalInfoResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortalInfoResponseModel struct{}"
	}

	return strings.Join([]string{"UpdatePortalInfoResponseModel", string(data)}, " ")
}
