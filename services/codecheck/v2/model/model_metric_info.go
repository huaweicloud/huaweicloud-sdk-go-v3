package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 度量项
type MetricInfo struct {

	// 代码规模
	CodeSize *string `json:"code_size,omitempty" xml:"code_size"`

	// 原始代码行数
	RawLines *string `json:"raw_lines,omitempty" xml:"raw_lines"`

	// 函数总数
	MethodsTotal *string `json:"methods_total,omitempty" xml:"methods_total"`

	// 圈复杂度总数
	CyclomaticComplexityTotal *string `json:"cyclomatic_complexity_total,omitempty" xml:"cyclomatic_complexity_total"`

	// 平均圈复杂度
	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty" xml:"cyclomatic_complexity_per_method"`

	// 最大圈复杂度
	MaximumCyclomaticComplexity *string `json:"maximum_cyclomatic_complexity,omitempty" xml:"maximum_cyclomatic_complexity"`

	// 超大圈复杂度数
	HugeCyclomaticComplexityTotal *string `json:"huge_cyclomatic_complexity_total,omitempty" xml:"huge_cyclomatic_complexity_total"`

	// 超大圈复杂度比例
	HugeCyclomaticComplexityRatio *string `json:"huge_cyclomatic_complexity_ratio,omitempty" xml:"huge_cyclomatic_complexity_ratio"`

	// cca圈复杂度总数
	CcaCyclomaticComplexityTotal *string `json:"cca_cyclomatic_complexity_total,omitempty" xml:"cca_cyclomatic_complexity_total"`

	// cca平均圈复杂度
	CcaCyclomaticComplexityPerMethod *string `json:"cca_cyclomatic_complexity_per_method,omitempty" xml:"cca_cyclomatic_complexity_per_method"`

	// cca最大圈复杂度
	MaximumCcaCyclomaticComplexity *string `json:"maximum_cca_cyclomatic_complexity,omitempty" xml:"maximum_cca_cyclomatic_complexity"`

	// 超大圈复杂度函数总数
	HugeCcaCyclomaticComplexityTotal *string `json:"huge_cca_cyclomatic_complexity_total,omitempty" xml:"huge_cca_cyclomatic_complexity_total"`

	// 圈复杂度满足度
	CyclomaticComplexityAdequacy *string `json:"cyclomatic_complexity_adequacy,omitempty" xml:"cyclomatic_complexity_adequacy"`

	// 最大深度
	MaximumDepth *string `json:"maximum_depth,omitempty" xml:"maximum_depth"`

	// 超大深度数
	HugeDepthTotal *string `json:"huge_depth_total,omitempty" xml:"huge_depth_total"`

	// 超大深度占比
	HugeDepthRatio *string `json:"huge_depth_ratio,omitempty" xml:"huge_depth_ratio"`

	// 函数总行数
	MethodLines *string `json:"method_lines,omitempty" xml:"method_lines"`

	// 函数平均代码行
	LinesPerMethod *string `json:"lines_per_method,omitempty" xml:"lines_per_method"`

	// 超大函数数
	HugeMethodTotal *string `json:"huge_method_total,omitempty" xml:"huge_method_total"`

	// 超大函数占比
	HugeMethodRatio *string `json:"huge_method_ratio,omitempty" xml:"huge_method_ratio"`

	// 文件总数
	FilesTotal *string `json:"files_total,omitempty" xml:"files_total"`

	// 目录总数
	FoldersTotal *string `json:"folders_total,omitempty" xml:"folders_total"`

	// 文件平均代码行
	LinesPerFile *string `json:"lines_per_file,omitempty" xml:"lines_per_file"`

	// 超大头文件数
	HugeHeaderfileTotal *string `json:"huge_headerfile_total,omitempty" xml:"huge_headerfile_total"`

	// 超大头文件占比
	HugeHeaderfileRatio *string `json:"huge_headerfile_ratio,omitempty" xml:"huge_headerfile_ratio"`

	// 超大源文件数
	HugeNonHeaderfileTotal *string `json:"huge_non_headerfile_total,omitempty" xml:"huge_non_headerfile_total"`

	// 超大源文件占比
	HugeNonHeaderfileRatio *string `json:"huge_non_headerfile_ratio,omitempty" xml:"huge_non_headerfile_ratio"`

	// 超大目录数
	HugeFolderTotal *string `json:"huge_folder_total,omitempty" xml:"huge_folder_total"`

	// 超大目录占比
	HugeFolderRatio *string `json:"huge_folder_ratio,omitempty" xml:"huge_folder_ratio"`

	// 重复文件数
	FileDuplicationTotal *string `json:"file_duplication_total,omitempty" xml:"file_duplication_total"`

	// 文件重复率
	FileDuplicationRatio *string `json:"file_duplication_ratio,omitempty" xml:"file_duplication_ratio"`

	// 重复源文件数
	NonHfileDuplicationTotal *string `json:"non_hfile_duplication_total,omitempty" xml:"non_hfile_duplication_total"`

	// 源文件重复率
	NonHfileDuplicationRatio *string `json:"non_hfile_duplication_ratio,omitempty" xml:"non_hfile_duplication_ratio"`

	// 代码重复数
	CodeDuplicationTotal *string `json:"code_duplication_total,omitempty" xml:"code_duplication_total"`

	// 代码重复率
	CodeDuplicationRatio *string `json:"code_duplication_ratio,omitempty" xml:"code_duplication_ratio"`

	// 源文件代码重复数
	NonHfileCodeDuplicationTotal *string `json:"non_hfile_code_duplication_total,omitempty" xml:"non_hfile_code_duplication_total"`

	// 源文件代码重复率
	NonHfileCodeDuplicationRatio *string `json:"non_hfile_code_duplication_ratio,omitempty" xml:"non_hfile_code_duplication_ratio"`

	// 危险函数总数
	UnsafeFunctionsTotal *string `json:"unsafe_functions_total,omitempty" xml:"unsafe_functions_total"`

	// 危险函数密度
	UnsafeFunctionsKloc *string `json:"unsafe_functions_kloc,omitempty" xml:"unsafe_functions_kloc"`

	// 冗余代码数
	RedundantCodeTotal *string `json:"redundant_code_total,omitempty" xml:"redundant_code_total"`

	// 冗余代码块密度
	RedundantCodeKloc *string `json:"redundant_code_kloc,omitempty" xml:"redundant_code_kloc"`

	// 抑制告警数
	WarningSuppressionTotal *string `json:"warning_suppression_total,omitempty" xml:"warning_suppression_total"`

	// 抑制告警密度
	WarningSuppressionKloc *string `json:"warning_suppression_kloc,omitempty" xml:"warning_suppression_kloc"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
