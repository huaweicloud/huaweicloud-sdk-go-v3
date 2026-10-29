package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricInfo struct {

	// **参数解释：** 有效代码行数，不包含空行和注释行。 **取值范围：** 不涉及。
	CodeSize *string `json:"code_size,omitempty"`

	// **参数解释：** 原始代码行数，包含空行和注释行。 **取值范围：** 不涉及。
	RawLines *string `json:"raw_lines,omitempty"`

	// **参数解释：** 函数总数。 **取值范围：** 不涉及。
	MethodsTotal *string `json:"methods_total,omitempty"`

	// **参数解释：** 圈复杂度总数。 **取值范围：** 不涉及。
	CyclomaticComplexityTotal *string `json:"cyclomatic_complexity_total,omitempty"`

	// **参数解释：** 平均圈复杂度。 **取值范围：** 不涉及。
	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty"`

	// **参数解释：** 最大圈复杂度。 **取值范围：** 不涉及。
	MaximumCyclomaticComplexity *string `json:"maximum_cyclomatic_complexity,omitempty"`

	// **参数解释：** 超大圈复杂度函数总数。 **取值范围：** 不涉及。
	HugeCyclomaticComplexityTotal *string `json:"huge_cyclomatic_complexity_total,omitempty"`

	// **参数解释：** 超大圈复杂度函数占比。 **取值范围：** 不涉及。
	HugeCyclomaticComplexityRatio *string `json:"huge_cyclomatic_complexity_ratio,omitempty"`

	// **参数解释：** CCA圈复杂度总数。 **取值范围：** 不涉及。
	CcaCyclomaticComplexityTotal *string `json:"cca_cyclomatic_complexity_total,omitempty"`

	// **参数解释：** CCA平均圈复杂度。 **取值范围：** 不涉及。
	CcaCyclomaticComplexityPerMethod *string `json:"cca_cyclomatic_complexity_per_method,omitempty"`

	// **参数解释：** CCA最大圈复杂度。 **取值范围：** 不涉及。
	MaximumCcaCyclomaticComplexity *string `json:"maximum_cca_cyclomatic_complexity,omitempty"`

	// **参数解释：** 超大CCA圈复杂度函数总数。 **取值范围：** 不涉及。
	HugeCcaCyclomaticComplexityTotal *string `json:"huge_cca_cyclomatic_complexity_total,omitempty"`

	// **参数解释：** 圈复杂度阈值遵从度。 **取值范围：** 不涉及。
	CyclomaticComplexityAdequacy *string `json:"cyclomatic_complexity_adequacy,omitempty"`

	// **参数解释：** 最大深度。 **取值范围：** 不涉及。
	MaximumDepth *string `json:"maximum_depth,omitempty"`

	// **参数解释：** 超大深度个数。 **取值范围：** 不涉及。
	HugeDepthTotal *string `json:"huge_depth_total,omitempty"`

	// **参数解释：** 超大深度占比。 **取值范围：** 不涉及。
	HugeDepthRatio *string `json:"huge_depth_ratio,omitempty"`

	// **参数解释：** 函数总行数。 **取值范围：** 不涉及。
	MethodLines *string `json:"method_lines,omitempty"`

	// **参数解释：** 平均函数代码行。 **取值范围：** 不涉及。
	LinesPerMethod *string `json:"lines_per_method,omitempty"`

	// **参数解释：** 超大代码行函数数。 **取值范围：** 不涉及。
	HugeMethodTotal *string `json:"huge_method_total,omitempty"`

	// **参数解释：** 超大代码行函数占比。 **取值范围：** 不涉及。
	HugeMethodRatio *string `json:"huge_method_ratio,omitempty"`

	// **参数解释：** 文件总数。 **取值范围：** 不涉及。
	FilesTotal *string `json:"files_total,omitempty"`

	// **参数解释：** 目录总数。 **取值范围：** 不涉及。
	FoldersTotal *string `json:"folders_total,omitempty"`

	// **参数解释：** 平均文件代码行。 **取值范围：** 不涉及。
	LinesPerFile *string `json:"lines_per_file,omitempty"`

	// **参数解释：** 超大头文件数。 **取值范围：** 不涉及。
	HugeHeaderfileTotal *string `json:"huge_headerfile_total,omitempty"`

	// **参数解释：** 超大头文件占比。 **取值范围：** 不涉及。
	HugeHeaderfileRatio *string `json:"huge_headerfile_ratio,omitempty"`

	// **参数解释：** 超大源文件数。 **取值范围：** 不涉及。
	HugeNonHeaderfileTotal *string `json:"huge_non_headerfile_total,omitempty"`

	// **参数解释：** 超大源文件占比。 **取值范围：** 不涉及。
	HugeNonHeaderfileRatio *string `json:"huge_non_headerfile_ratio,omitempty"`

	// **参数解释：** 超大目录数。 **取值范围：** 不涉及。
	HugeFolderTotal *string `json:"huge_folder_total,omitempty"`

	// **参数解释：** 超大目录占比。 **取值范围：** 不涉及。
	HugeFolderRatio *string `json:"huge_folder_ratio,omitempty"`

	// **参数解释：** 重复文件数。 **取值范围：** 不涉及。
	FileDuplicationTotal *string `json:"file_duplication_total,omitempty"`

	// **参数解释：** 文件重复率。 **取值范围：** 不涉及。
	FileDuplicationRatio *string `json:"file_duplication_ratio,omitempty"`

	// **参数解释：** 重复源文件数。 **取值范围：** 不涉及。
	NonHfileDuplicationTotal *string `json:"non_hfile_duplication_total,omitempty"`

	// **参数解释：** 源文件重复率。 **取值范围：** 不涉及。
	NonHfileDuplicationRatio *string `json:"non_hfile_duplication_ratio,omitempty"`

	// **参数解释：** 代码重复数。 **取值范围：** 不涉及。
	CodeDuplicationTotal *string `json:"code_duplication_total,omitempty"`

	// **参数解释：** 代码重复率。 **取值范围：** 不涉及。
	CodeDuplicationRatio *string `json:"code_duplication_ratio,omitempty"`

	// **参数解释：** 源文件代码重复数。 **取值范围：** 不涉及。
	NonHfileCodeDuplicationTotal *string `json:"non_hfile_code_duplication_total,omitempty"`

	// **参数解释：** 源文件代码重复率。 **取值范围：** 不涉及。
	NonHfileCodeDuplicationRatio *string `json:"non_hfile_code_duplication_ratio,omitempty"`

	// **参数解释：** 危险函数总数。 **取值范围：** 不涉及。
	UnsafeFunctionsTotal *string `json:"unsafe_functions_total,omitempty"`

	// **参数解释：** 危险函数密度。 **取值范围：** 不涉及。
	UnsafeFunctionsKloc *string `json:"unsafe_functions_kloc,omitempty"`

	// **参数解释：** 冗余代码数。 **取值范围：** 不涉及。
	RedundantCodeTotal *string `json:"redundant_code_total,omitempty"`

	// **参数解释：** 冗余代码密度。 **取值范围：** 不涉及。
	RedundantCodeKloc *string `json:"redundant_code_kloc,omitempty"`

	// **参数解释：** 抑制告警数。 **取值范围：** 不涉及。
	WarningSuppressionTotal *string `json:"warning_suppression_total,omitempty"`

	// **参数解释：** 抑制告警密度。 **取值范围：** 不涉及。
	WarningSuppressionKloc *string `json:"warning_suppression_kloc,omitempty"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
