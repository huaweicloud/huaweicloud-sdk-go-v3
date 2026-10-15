package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSparkSqlResponse Response Object
type CancelSparkSqlResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSparkSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkSqlResponse struct{}"
	}

	return strings.Join([]string{"CancelSparkSqlResponse", string(data)}, " ")
}
