package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCicdConfigurationsRequestBody 需要删除的CI/CD标识列表
type DeleteCicdConfigurationsRequestBody struct {

	// CI/CD标识列表
	CicdIdList []string `json:"cicd_id_list"`
}

func (o DeleteCicdConfigurationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCicdConfigurationsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteCicdConfigurationsRequestBody", string(data)}, " ")
}
