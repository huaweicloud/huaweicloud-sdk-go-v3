package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPictureModelingJobsRequest Request Object
type ListPictureModelingJobsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。格式为YYYYMMDD'T'HHMMSS'Z'
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 任务状态，默认所有状态，可多个状态查询，使用英文逗号分隔。 * WAITING：等待任务调度 * PROCESSING：正在处理 * PARTIAL_SUCCEED: 部分成功（模型生成，截图失败） * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State *string `json:"state,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 升序还是降序，默认升序。 * asc：升序 * desc：降序
	SortDir *ListPictureModelingJobsRequestSortDir `json:"sort_dir,omitempty"`

	// 过滤创建时间<=输入时间的记录。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"''。
	CreateUntil *string `json:"create_until,omitempty"`

	// 过滤创建时间>=输入时间的记录。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"''。
	CreateSince *string `json:"create_since,omitempty"`
}

func (o ListPictureModelingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPictureModelingJobsRequest struct{}"
	}

	return strings.Join([]string{"ListPictureModelingJobsRequest", string(data)}, " ")
}

type ListPictureModelingJobsRequestSortDir struct {
	value string
}

type ListPictureModelingJobsRequestSortDirEnum struct {
	ASC  ListPictureModelingJobsRequestSortDir
	DESC ListPictureModelingJobsRequestSortDir
}

func GetListPictureModelingJobsRequestSortDirEnum() ListPictureModelingJobsRequestSortDirEnum {
	return ListPictureModelingJobsRequestSortDirEnum{
		ASC: ListPictureModelingJobsRequestSortDir{
			value: "asc",
		},
		DESC: ListPictureModelingJobsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPictureModelingJobsRequestSortDir) Value() string {
	return c.value
}

func (c ListPictureModelingJobsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPictureModelingJobsRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
