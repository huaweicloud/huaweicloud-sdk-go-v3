package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowHdfsFileListRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。

	ClusterId string `json:"cluster_id"`
	// 文件目录，比如访问“/tmp/test”目录列表，此处必须是目录，整体URI为  /v2/{project_id}/clusters/{cluster_id}/files?path=%2Ftmp%2Ftest  单层目录要遵循以下规则：  - 不能为空 - 不能以\".\"开头或结尾 - 不能包括下列符号 : :*?\"<>|\\;&,'`!{}[]$%+ - 不能超过255个字节

	Path string `json:"path"`
	// 分页参数，表示从该偏移量开始查询文件列表，默认值为1。

	Offset *string `json:"offset,omitempty"`
	// 分页参数，列表当前分页的数量限制，默认为100，最大1000。

	Limit *string `json:"limit,omitempty"`
	// 列表排序按该属性排序。缺省值：path_suffix  - path_suffix：文件或目录名称 - length：文件大小 - modification_time：修改时间

	SortKey *ShowHdfsFileListRequestSortKey `json:"sort_key,omitempty"`
	// 列表排序方式，desc为降序，asc为升序，默认值为desc。

	Order *ShowHdfsFileListRequestOrder `json:"order,omitempty"`
}

func (o ShowHdfsFileListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdfsFileListRequest struct{}"
	}

	return strings.Join([]string{"ShowHdfsFileListRequest", string(data)}, " ")
}

type ShowHdfsFileListRequestSortKey struct {
	value string
}

type ShowHdfsFileListRequestSortKeyEnum struct {
	PATH_SUFFIX       ShowHdfsFileListRequestSortKey
	LENGTH            ShowHdfsFileListRequestSortKey
	MODIFICATION_TIME ShowHdfsFileListRequestSortKey
}

func GetShowHdfsFileListRequestSortKeyEnum() ShowHdfsFileListRequestSortKeyEnum {
	return ShowHdfsFileListRequestSortKeyEnum{
		PATH_SUFFIX: ShowHdfsFileListRequestSortKey{
			value: "path_suffix：文件或目录名称",
		},
		LENGTH: ShowHdfsFileListRequestSortKey{
			value: "length：文件大小",
		},
		MODIFICATION_TIME: ShowHdfsFileListRequestSortKey{
			value: "modification_time：修改时间",
		},
	}
}

func (c ShowHdfsFileListRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHdfsFileListRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowHdfsFileListRequestOrder struct {
	value string
}

type ShowHdfsFileListRequestOrderEnum struct {
	DESC ShowHdfsFileListRequestOrder
	ASC  ShowHdfsFileListRequestOrder
}

func GetShowHdfsFileListRequestOrderEnum() ShowHdfsFileListRequestOrderEnum {
	return ShowHdfsFileListRequestOrderEnum{
		DESC: ShowHdfsFileListRequestOrder{
			value: "desc",
		},
		ASC: ShowHdfsFileListRequestOrder{
			value: "asc",
		},
	}
}

func (c ShowHdfsFileListRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHdfsFileListRequestOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
