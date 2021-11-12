package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsurantItem struct {
	InsurantName *InsurancePolicyDetail `json:"insurant_name,omitempty"`

	InsurantSex *InsurancePolicyDetail `json:"insurant_sex,omitempty"`

	InsurantBirthday *InsurancePolicyDetail `json:"insurant_birthday,omitempty"`

	InsurantIdType *InsurancePolicyDetail `json:"insurant_id_type,omitempty"`

	InsurantIdNumber *InsurancePolicyDetail `json:"insurant_id_number,omitempty"`
}

func (o InsurantItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsurantItem struct{}"
	}

	return strings.Join([]string{"InsurantItem", string(data)}, " ")
}
