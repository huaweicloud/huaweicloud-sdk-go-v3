package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BillingResult 批价基础结果。
type BillingResult struct {
}

func (o BillingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingResult struct{}"
	}

	return strings.Join([]string{"BillingResult", string(data)}, " ")
}
