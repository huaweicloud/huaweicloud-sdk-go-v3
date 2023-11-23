package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDriverReq struct {

	// jdbc驱动文件名称，name的长度5-64，结尾以.jar结尾。
	DriverName string `json:"driver_name"`
}

func (o UpdateDriverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDriverReq struct{}"
	}

	return strings.Join([]string{"UpdateDriverReq", string(data)}, " ")
}
