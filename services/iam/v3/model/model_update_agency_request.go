package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAgencyRequest struct {
	// 待修改的委托ID，获取方式请参见：[获取委托名、委托ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	AgencyId string `json:"agency_id"`

	Body *UpdateAgencyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyRequest", string(data)}, " ")
}
