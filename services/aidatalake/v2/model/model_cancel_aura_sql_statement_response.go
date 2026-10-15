package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAuraSqlStatementResponse Response Object
type CancelAuraSqlStatementResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAuraSqlStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAuraSqlStatementResponse struct{}"
	}

	return strings.Join([]string{"CancelAuraSqlStatementResponse", string(data)}, " ")
}
