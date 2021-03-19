package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceUsagesResponse struct {
	// 套餐包使用量信息，具体请参见表2。

	PackageUsageInfos *[]PackageUsageInfo `json:"package_usage_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListResourceUsagesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceUsagesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceUsagesResponse", string(data)}, " ")
}
