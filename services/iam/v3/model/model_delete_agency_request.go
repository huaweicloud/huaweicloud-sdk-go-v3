package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAgencyRequest struct {
	// 待删除的委托ID，获取方式请参见：[获取委托名、委托ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	AgencyId string `json:"agency_id"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}
