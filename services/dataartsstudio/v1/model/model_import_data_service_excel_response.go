package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataServiceExcelResponse Response Object
type ImportDataServiceExcelResponse struct {

	// 返回的数据信息。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportDataServiceExcelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataServiceExcelResponse struct{}"
	}

	return strings.Join([]string{"ImportDataServiceExcelResponse", string(data)}, " ")
}
