/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AndroidUnionChannelPackagesAddResponseData struct {
	AppAndroidChannelPackageId string        `json:"app_android_channel_package_id,omitempty"`
	PackageName                string        `json:"package_name,omitempty"`
	PackageStatus              PackageStatus `json:"package_status,omitempty"`
}