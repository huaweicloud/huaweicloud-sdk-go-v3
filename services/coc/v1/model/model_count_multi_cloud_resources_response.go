package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMultiCloudResourcesResponse Response Object
type CountMultiCloudResourcesResponse struct {

	// **参数解释：** 云资源数量。 **取值范围：** 不涉及。
	Data           *int64 `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountMultiCloudResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMultiCloudResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountMultiCloudResourcesResponse", string(data)}, " ")
}
