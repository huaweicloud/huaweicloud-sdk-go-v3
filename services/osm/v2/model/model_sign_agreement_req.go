package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignAgreementReq struct {

	// 关联id
	RelationId *string `json:"relation_id,omitempty"`
}

func (o SignAgreementReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignAgreementReq struct{}"
	}

	return strings.Join([]string{"SignAgreementReq", string(data)}, " ")
}
