package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// API版本详细信息列表。
type ApiVersionResponse struct {

	// API版本号。
	Id string `json:"id" xml:"id"`

	// 对应API的链接信息,v3版本该字段为[]。
	Links []Links `json:"links" xml:"links"`

	// 版本状态。 取值为“CURRENT”，表示该版本目前已对外公布。
	Status string `json:"status" xml:"status"`

	// API版本的子版本信息。
	Version string `json:"version" xml:"version"`

	// API版本的最小版本号。
	MinVersion string `json:"min_version" xml:"min_version"`

	// 版本更新时间。格式为“yyyy-mm-dd Thh:mm:ssZ”。其中，T指某个时间的开始，Z指UTC时间。
	Updated string `json:"updated" xml:"updated"`
}

func (o ApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ApiVersionResponse", string(data)}, " ")
}
