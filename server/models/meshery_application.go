package models

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	isql "github.com/layer5io/meshery/server/internal/sql"
)

type ApplicationType string

type ApplicationTypeResponse struct {
	Type                ApplicationType `json:"application_type"`
	SupportedExtensions []string        `json:"supported_extensions"`
}

func GetApplicationTypes() (r []ApplicationTypeResponse) {
	r = append(r, ApplicationTypeResponse{
		Type:                HelmChart,
		SupportedExtensions: []string{".tgz"},
	},
		ApplicationTypeResponse{
			Type:                DockerCompose,
			SupportedExtensions: []string{".yaml", ".yml"},
		},
		ApplicationTypeResponse{
			Type:                K8sManifest,
			SupportedExtensions: []string{".yaml", ".yml"},
		})
	return
}

const (
	HelmChart     ApplicationType = "Helm Chart"
	DockerCompose ApplicationType = "Docker Compose"
	K8sManifest   ApplicationType = "Kubernetes Manifest"
)

// MesheryApplication represents the applications that needs to be saved
type MesheryApplication struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Name            string `json:"name,omitempty"`
	ApplicationFile string `json:"application_file"`
	// Meshery doesn't have the user id fields
	// but the remote provider is allowed to provide one
	UserID        *string        `json:"user_id" gorm:"-"`
	Location      isql.Map       `json:"location"`
	Type          sql.NullString `json:"type"`
	SourceContent []byte         `json:"source_content"`
	UpdatedAt     *time.Time     `json:"updated_at,omitempty"`
	CreatedAt     *time.Time     `json:"created_at,omitempty"`
}
