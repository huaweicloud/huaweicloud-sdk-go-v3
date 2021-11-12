package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Authentification struct {
	// 域名校验值名字。

	RecordName *string `json:"record_name,omitempty"`
	// 域名校验值类型。

	RecordType *string `json:"record_type,omitempty"`
	// 域名校验值。

	RecordValue *string `json:"record_value,omitempty"`
	// 校验值对应的域名。

	Domain *string `json:"domain,omitempty"`
}

func (o Authentification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authentification struct{}"
	}

	return strings.Join([]string{"Authentification", string(data)}, " ")
}
