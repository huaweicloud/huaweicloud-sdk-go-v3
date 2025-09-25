package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageWhiteListsRequestBodyQueryInfo 按照镜像查询条件指定镜像
type AddImageWhiteListsRequestBodyQueryInfo struct {

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image：私有镜像仓库。 - shared_image：共享镜像仓库。 - instance_image：企业仓库。 - harbor: harbor仓库。 - jfrog : jfrog仓库。  **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`
}

func (o AddImageWhiteListsRequestBodyQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageWhiteListsRequestBodyQueryInfo struct{}"
	}

	return strings.Join([]string{"AddImageWhiteListsRequestBodyQueryInfo", string(data)}, " ")
}
