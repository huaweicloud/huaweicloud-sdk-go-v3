package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// API版本详细信息列表。
type ApiVersion struct {
	// API版本号，如v3。

	Id string `json:"id"`
	// 版本状态。 取值“CURRENT”，表示该版本为主推版本。取值“SUPPORTED”，表示为老版本，但是现在还继续支持。 取值“DEPRECATED”，表示为废弃版本，存在后续删除的可能。

	Status string `json:"status"`
	// 版本发布时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指UTC时间。

	Updated string `json:"updated"`
	// API的微版本，如果不支持微版本，则为空

	Version string `json:"version"`
}

func (o ApiVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersion struct{}"
	}

	return strings.Join([]string{"ApiVersion", string(data)}, " ")
}
