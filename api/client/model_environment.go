/*
 * Merlin
 *
 * API Guide for accessing Merlin's model management, deployment, and serving functionalities
 *
 * API version: 0.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type Environment struct {
	Id                     int32            `json:"id,omitempty"`
	Name                   string           `json:"name"`
	Cluster                string           `json:"cluster,omitempty"`
	IsDefault              bool             `json:"is_default,omitempty"`
	Region                 string           `json:"region,omitempty"`
	GcpProject             string           `json:"gcp_project,omitempty"`
	DefaultResourceRequest *ResourceRequest `json:"default_resource_request,omitempty"`
	CreatedAt              time.Time        `json:"created_at,omitempty"`
	UpdatedAt              time.Time        `json:"updated_at,omitempty"`
}
