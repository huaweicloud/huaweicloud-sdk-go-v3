package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataMapLineageResponse Response Object
type ImportDataMapLineageResponse struct {

	// 血缘信息资产guids。
	Guids          *[]NodeLineageGuids `json:"guids,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ImportDataMapLineageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataMapLineageResponse struct{}"
	}

	return strings.Join([]string{"ImportDataMapLineageResponse", string(data)}, " ")
}
