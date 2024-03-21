package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HouseholdRegisterContent struct {

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 户主或与户主关系。
	HouseholderRelationship *string `json:"householder_relationship,omitempty"`

	// 曾用名。
	FormerName *string `json:"former_name,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 出生地。
	Birthplace *string `json:"birthplace,omitempty"`

	// 民族。
	Ethnicity *string `json:"ethnicity,omitempty"`

	// 籍贯。
	OriginPlace *string `json:"origin_place,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 本市(县)其他住址。
	OtherAddress *string `json:"other_address,omitempty"`

	// 宗教信仰。
	ReligiousBelief *string `json:"religious_belief,omitempty"`

	// 公民身份证件编号。
	IdCardNumber *string `json:"id_card_number,omitempty"`

	// 身高。
	Height *string `json:"height,omitempty"`

	// 血型。
	BloodType *string `json:"blood_type,omitempty"`

	// 文化程度。
	Education *string `json:"education,omitempty"`

	// 婚姻状况。
	MaritalStatus *string `json:"marital_status,omitempty"`

	// 兵役情况。
	MilitaryServiceStatus *string `json:"military_service_status,omitempty"`

	// 服务处所。
	WorkPlace *string `json:"work_place,omitempty"`

	// 职业。
	Occupation *string `json:"occupation,omitempty"`

	// 何时由何地迁来本市(县)。
	MigratedToCity *string `json:"migrated_to_city,omitempty"`

	// 何时由何地迁来本址。
	MigratedToAddress *string `json:"migrated_to_address,omitempty"`

	// 承办人签章。
	RegistrarSignatureSeal *string `json:"registrar_signature_seal,omitempty"`

	// 登记日期。
	RegistrationDate *string `json:"registration_date,omitempty"`

	// 户别。
	HouseholdType *string `json:"household_type,omitempty"`

	// 户号。
	HouseholdNumber *string `json:"household_number,omitempty"`

	// 户主姓名。当type参数为“首页”时，返回此参数。
	HouseholderName *string `json:"householder_name,omitempty"`

	// 社区。当type参数为“首页”时，返回此参数。
	Community *string `json:"community,omitempty"`

	// 住址。当type参数为“首页”时，返回此参数。
	Address *string `json:"address,omitempty"`

	// 签发日期。当type参数为“首页”时，返回此参数。
	IssueDate *string `json:"issue_date,omitempty"`

	// 户口登记机关。当type参数为“首页”时，返回此参数。
	PoliceStation *string `json:"police_station,omitempty"`
}

func (o HouseholdRegisterContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HouseholdRegisterContent struct{}"
	}

	return strings.Join([]string{"HouseholdRegisterContent", string(data)}, " ")
}
