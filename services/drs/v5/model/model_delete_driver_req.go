package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDriverReq struct {

	// jdbc驱动文件列表，列表长度1-20，driver_name的长度5-64，结尾以.jar结尾。
	DriverNames []string `json:"driver_names"`
}

func (o DeleteDriverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDriverReq struct{}"
	}

	return strings.Join([]string{"DeleteDriverReq", string(data)}, " ")
}
