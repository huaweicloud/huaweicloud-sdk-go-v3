package model

import (
	"encoding/json"

	"strings"
)

type ExtentionReqDataByNameAndId struct {
	// 被验证人的姓名。

	VerificationName string `json:"verification_name"`
	// 被验证人的身份证号码。

	VerificationId string `json:"verification_id"`
}

func (o ExtentionReqDataByNameAndId) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtentionReqDataByNameAndId struct{}"
	}

	return strings.Join([]string{"ExtentionReqDataByNameAndId", string(data)}, " ")
}
