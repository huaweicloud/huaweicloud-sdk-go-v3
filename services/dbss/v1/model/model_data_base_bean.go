package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataBaseBean struct {
	Database *DataBase `json:"database"`
}

func (o DataBaseBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBaseBean struct{}"
	}

	return strings.Join([]string{"DataBaseBean", string(data)}, " ")
}
