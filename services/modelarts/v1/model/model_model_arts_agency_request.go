package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModelArtsAgencyRequest struct {

	// 委托名称后缀。  长度不大于50位。  委托名称前缀固定为ma_agency。  如该字段为iam-user01，则创建出来的委托名称为ma_agency_iam-user01。  默认为空，表示创建名称为modelarts_agency的委托。
	AgencyNameSuffix *string `json:"agency_name_suffix,omitempty"`
}

func (o ModelArtsAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelArtsAgencyRequest struct{}"
	}

	return strings.Join([]string{"ModelArtsAgencyRequest", string(data)}, " ")
}
