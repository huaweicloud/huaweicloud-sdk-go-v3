package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFavoriteRequest Request Object
type ListFavoriteRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 收藏人名称列表。
	UserNameList *[]string `json:"user_name_list,omitempty"`

	// 资源类型列表。
	ResourceTypeList *[]string `json:"resource_type_list,omitempty"`

	// 收藏类型列表，支持MICROMOLECULE|PROTEIN。
	TypeList *[]string `json:"type_list,omitempty"`

	// 查询收藏信息的起始时间，UNIX时间戳，单位毫秒，不填时默认为当前时间。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询收藏信息的结束时间，UNIX时间戳，单位毫秒，不填时默认为当前时间。
	EndTime *int64 `json:"end_time,omitempty"`

	// 关键字，支持在display_info字段中内容的模糊搜索。
	KeyWord *string `json:"key_word,omitempty"`

	// 排序规则，目前默认时间降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则，目前默认按收藏时间降序，支持根据create_time|user_name|resource_name|resource_type排序。
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListFavoriteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFavoriteRequest struct{}"
	}

	return strings.Join([]string{"ListFavoriteRequest", string(data)}, " ")
}
