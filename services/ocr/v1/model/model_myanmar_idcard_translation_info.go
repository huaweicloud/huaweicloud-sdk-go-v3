package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarIdcardTranslationInfo struct {

	// 名字转译。仅当输入参数return_translation为true时，返回该字段。
	NameTranslation *string `json:"name_translation,omitempty"`

	// 父亲名字的转译。仅当输入参数return_translation为true时，返回该字段。
	FatherNameTranslation *string `json:"father_name_translation,omitempty"`

	// 身份证号码转译。仅当输入参数return_translation为true时，返回该字段。
	NrcIdTranslation *string `json:"nrc_id_translation,omitempty"`

	// 出生日期转译。仅当输入参数return_translation为true时，返回该字段。
	BirthTranslation *string `json:"birth_translation,omitempty"`
}

func (o MyanmarIdcardTranslationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarIdcardTranslationInfo struct{}"
	}

	return strings.Join([]string{"MyanmarIdcardTranslationInfo", string(data)}, " ")
}
