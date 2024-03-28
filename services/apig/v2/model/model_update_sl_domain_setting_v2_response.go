package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlDomainSettingV2Response Response Object
type UpdateSlDomainSettingV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlDomainSettingV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlDomainSettingV2Response struct{}"
	}

	return strings.Join([]string{"UpdateSlDomainSettingV2Response", string(data)}, " ")
}
