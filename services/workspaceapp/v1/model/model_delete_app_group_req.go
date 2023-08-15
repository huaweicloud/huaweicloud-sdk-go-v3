package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppGroupReq 批量删除应用
type DeleteAppGroupReq struct {

	// 应用组ID,单次最多20个
	Ids []string `json:"ids"`
}

func (o DeleteAppGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppGroupReq struct{}"
	}

	return strings.Join([]string{"DeleteAppGroupReq", string(data)}, " ")
}
