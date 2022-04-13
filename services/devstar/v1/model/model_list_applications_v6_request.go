package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListApplicationsV6Request struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *ListApplicationsV6RequestXLanguage `json:"X-Language,omitempty"`
	// 是否查询我关注的应用

	Attention *bool `json:"attention,omitempty"`
	// 区域id，从控制台获取方法请参见: [获取区域ID](https://console.huaweicloud.com/iam/?region=cn-north-1&locale=zh-cn#/iam/projects)

	RegionId *string `json:"region_id,omitempty"`
	// 搜索关键字,支持按名称和描述搜索，默认null

	Keyword *string `json:"keyword,omitempty"`
	// 所属DevCloud项目id，从 [项目列表接口](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DevStar&api=ListProjectsV4) 查询。

	ProjectId *string `json:"project_id,omitempty"`
	// 主题id,场景或者部署方式分类id

	TopicId *string `json:"topic_id,omitempty"`
	// 是否查询由我创建

	IsCreatedBySelf *bool `json:"is_created_by_self,omitempty"`
	// 排序字段, name：应用名称,created_at:创建时间,updated_at:更新时间

	SortKey *[]ListApplicationsV6RequestSortKey `json:"sort_key,omitempty"`
	// 排序方式, desc:降序, asc:升序

	SortDir *[]ListApplicationsV6RequestSortDir `json:"sort_dir,omitempty"`
	// 每页显示的条目数量,默认10

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询,默认0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListApplicationsV6Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsV6Request struct{}"
	}

	return strings.Join([]string{"ListApplicationsV6Request", string(data)}, " ")
}

type ListApplicationsV6RequestXLanguage struct {
	value string
}

type ListApplicationsV6RequestXLanguageEnum struct {
	ZH_CN ListApplicationsV6RequestXLanguage
	EN_US ListApplicationsV6RequestXLanguage
}

func GetListApplicationsV6RequestXLanguageEnum() ListApplicationsV6RequestXLanguageEnum {
	return ListApplicationsV6RequestXLanguageEnum{
		ZH_CN: ListApplicationsV6RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListApplicationsV6RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListApplicationsV6RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationsV6RequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListApplicationsV6RequestSortKey struct {
	value string
}

type ListApplicationsV6RequestSortKeyEnum struct {
	NAME       ListApplicationsV6RequestSortKey
	CREATED_AT ListApplicationsV6RequestSortKey
	UPDATED_AT ListApplicationsV6RequestSortKey
}

func GetListApplicationsV6RequestSortKeyEnum() ListApplicationsV6RequestSortKeyEnum {
	return ListApplicationsV6RequestSortKeyEnum{
		NAME: ListApplicationsV6RequestSortKey{
			value: "name",
		},
		CREATED_AT: ListApplicationsV6RequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListApplicationsV6RequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListApplicationsV6RequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationsV6RequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListApplicationsV6RequestSortDir struct {
	value string
}

type ListApplicationsV6RequestSortDirEnum struct {
	DESC ListApplicationsV6RequestSortDir
	ASC  ListApplicationsV6RequestSortDir
}

func GetListApplicationsV6RequestSortDirEnum() ListApplicationsV6RequestSortDirEnum {
	return ListApplicationsV6RequestSortDirEnum{
		DESC: ListApplicationsV6RequestSortDir{
			value: "desc",
		},
		ASC: ListApplicationsV6RequestSortDir{
			value: "asc",
		},
	}
}

func (c ListApplicationsV6RequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationsV6RequestSortDir) UnmarshalJSON(b []byte) error {
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
