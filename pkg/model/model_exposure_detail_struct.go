/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 诊断详情-曝光明细
type ExposureDetailStruct struct {
	EffectData                 ExposureEffectDataStruct                 `json:"effect_data,omitempty"`
	EffectDataTrends           []ExposureEffectDataTrendsItem           `json:"effect_data_trends,omitempty"`
	TargetingLabelContribution ExposureTargetingLabelContributionStruct `json:"targeting_label_contribution,omitempty"`
}
