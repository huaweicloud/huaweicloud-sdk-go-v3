package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 度量项
type MetricInfo struct {

	// 代码规模
	CodeSize *string `json:"code_size,omitempty"`

	// 原始代码行数
	RawLines *string `json:"raw_lines,omitempty"`

	// 函数总数
	MethodsTotal *string `json:"methods_total,omitempty"`

	// 圈复杂度总数
	CyclomaticComplexityTotal *string `json:"cyclomatic_complexity_total,omitempty"`

	// 平均圈复杂度
	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty"`

	// 最大圈复杂度
	MaximumCyclomaticComplexity *string `json:"maximum_cyclomatic_complexity,omitempty"`

	// 超大圈复杂度数
	HugeCyclomaticComplexityTotal *string `json:"huge_cyclomatic_complexity_total,omitempty"`

	// 超大圈复杂度比例
	HugeCyclomaticComplexityRatio *string `json:"huge_cyclomatic_complexity_ratio,omitempty"`

	// cca圈复杂度总数
	CcaCyclomaticComplexityTotal *string `json:"cca_cyclomatic_complexity_total,omitempty"`

	// cca平均圈复杂度
	CcaCyclomaticComplexityPerMethod *string `json:"cca_cyclomatic_complexity_per_method,omitempty"`

	// cca最大圈复杂度
	MaximumCcaCyclomaticComplexity *string `json:"maximum_cca_cyclomatic_complexity,omitempty"`

	// 超大圈复杂度函数总数
	HugeCcaCyclomaticComplexityTotal *string `json:"huge_cca_cyclomatic_complexity_total,omitempty"`

	// 圈复杂度满足度
	CyclomaticComplexityAdequacy *string `json:"cyclomatic_complexity_adequacy,omitempty"`

	// 最大深度
	MaximumDepth *string `json:"maximum_depth,omitempty"`

	// 超大深度数
	HugeDepthTotal *string `json:"huge_depth_total,omitempty"`

	// 超大深度占比
	HugeDepthRatio *string `json:"huge_depth_ratio,omitempty"`

	// 函数总行数
	MethodLines *string `json:"method_lines,omitempty"`

	// 函数平均代码行
	LinesPerMethod *string `json:"lines_per_method,omitempty"`

	// 超大函数数
	HugeMethodTotal *string `json:"huge_method_total,omitempty"`

	// 超大函数占比
	HugeMethodRatio *string `json:"huge_method_ratio,omitempty"`

	// 文件总数
	FilesTotal *string `json:"files_total,omitempty"`

	// 目录总数
	FoldersTotal *string `json:"folders_total,omitempty"`

	// 文件平均代码行
	LinesPerFile *string `json:"lines_per_file,omitempty"`

	// 超大头文件数
	HugeHeaderfileTotal *string `json:"huge_headerfile_total,omitempty"`

	// 超大头文件占比
	HugeHeaderfileRatio *string `json:"huge_headerfile_ratio,omitempty"`

	// 超大源文件数
	HugeNonHeaderfileTotal *string `json:"huge_non_headerfile_total,omitempty"`

	// 超大源文件占比
	HugeNonHeaderfileRatio *string `json:"huge_non_headerfile_ratio,omitempty"`

	// 超大目录数
	HugeFolderTotal *string `json:"huge_folder_total,omitempty"`

	// 超大目录占比
	HugeFolderRatio *string `json:"huge_folder_ratio,omitempty"`

	// 重复文件数
	FileDuplicationTotal *string `json:"file_duplication_total,omitempty"`

	// 文件重复率
	FileDuplicationRatio *string `json:"file_duplication_ratio,omitempty"`

	// 重复源文件数
	NonHfileDuplicationTotal *string `json:"non_hfile_duplication_total,omitempty"`

	// 源文件重复率
	NonHfileDuplicationRatio *string `json:"non_hfile_duplication_ratio,omitempty"`

	// 代码重复数
	CodeDuplicationTotal *string `json:"code_duplication_total,omitempty"`

	// 代码重复率
	CodeDuplicationRatio *string `json:"code_duplication_ratio,omitempty"`

	// 源文件代码重复数
	NonHfileCodeDuplicationTotal *string `json:"non_hfile_code_duplication_total,omitempty"`

	// 源文件代码重复率
	NonHfileCodeDuplicationRatio *string `json:"non_hfile_code_duplication_ratio,omitempty"`

	// 危险函数总数
	UnsafeFunctionsTotal *string `json:"unsafe_functions_total,omitempty"`

	// 危险函数密度
	UnsafeFunctionsKloc *string `json:"unsafe_functions_kloc,omitempty"`

	// 冗余代码数
	RedundantCodeTotal *string `json:"redundant_code_total,omitempty"`

	// 冗余代码块密度
	RedundantCodeKloc *string `json:"redundant_code_kloc,omitempty"`

	// 抑制告警数
	WarningSuppressionTotal *string `json:"warning_suppression_total,omitempty"`

	// 抑制告警密度
	WarningSuppressionKloc *string `json:"warning_suppression_kloc,omitempty"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
