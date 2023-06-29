package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenBulkClassifications struct {

	// 数据资产list
	Guids []string `json:"guids"`

	Classification *OpenClassification `json:"classification"`
}

func (o OpenBulkClassifications) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenBulkClassifications struct{}"
	}

	return strings.Join([]string{"OpenBulkClassifications", string(data)}, " ")
}
