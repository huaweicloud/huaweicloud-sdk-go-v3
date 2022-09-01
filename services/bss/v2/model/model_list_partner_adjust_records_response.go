package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPartnerAdjustRecordsResponse struct {

	// 调账记录列表。 具体请参见表2。
	Records *[]AdjustRecordV3 `json:"records,omitempty" xml:"records"`

	// 返回总条数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPartnerAdjustRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerAdjustRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPartnerAdjustRecordsResponse", string(data)}, " ")
}
