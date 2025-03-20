package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccountSummaryV5Request Request Object
type GetAccountSummaryV5Request struct {
}

func (o GetAccountSummaryV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccountSummaryV5Request struct{}"
	}

	return strings.Join([]string{"GetAccountSummaryV5Request", string(data)}, " ")
}
