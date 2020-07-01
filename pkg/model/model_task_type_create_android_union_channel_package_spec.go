/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创建广告渠道包接口任务所需条件
type TaskTypeCreateAndroidUnionChannelPackageSpec struct {
	AndroidUnionAppId              int64                                  `json:"android_union_app_id,omitempty"`
	AndroidUnionChannelPackageSpec []CreateAndroidUnionChannelPackageSpec `json:"android_union_channel_package_spec,omitempty"`
}