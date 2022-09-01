package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsurantItem struct {
	InsurantName *InsurancePolicyDetail `json:"insurant_name,omitempty" xml:"insurant_name"`

	InsurantSex *InsurancePolicyDetail `json:"insurant_sex,omitempty" xml:"insurant_sex"`

	InsurantBirthday *InsurancePolicyDetail `json:"insurant_birthday,omitempty" xml:"insurant_birthday"`

	InsurantIdType *InsurancePolicyDetail `json:"insurant_id_type,omitempty" xml:"insurant_id_type"`

	InsurantIdNumber *InsurancePolicyDetail `json:"insurant_id_number,omitempty" xml:"insurant_id_number"`
}

func (o InsurantItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsurantItem struct{}"
	}

	return strings.Join([]string{"InsurantItem", string(data)}, " ")
}
