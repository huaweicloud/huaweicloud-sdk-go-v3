package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSearchCriterias struct {

	// 企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 快速查询id
	Id string `json:"id"`
}

func (o DeleteSearchCriterias) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSearchCriterias struct{}"
	}

	return strings.Join([]string{"DeleteSearchCriterias", string(data)}, " ")
}
