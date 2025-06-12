package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserOverPackageQuotaRequest Request Object
type ShowUserOverPackageQuotaRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`
}

func (o ShowUserOverPackageQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserOverPackageQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowUserOverPackageQuotaRequest", string(data)}, " ")
}
