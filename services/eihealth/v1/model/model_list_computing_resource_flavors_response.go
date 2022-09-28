package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListComputingResourceFlavorsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 规格列表
	Flavors        *[]ComputingResourceFlavorsRsp `json:"flavors,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListComputingResourceFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourceFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListComputingResourceFlavorsResponse", string(data)}, " ")
}
