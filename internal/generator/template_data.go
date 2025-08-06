package generator

import (
	"fmt"
	"strings"
	"time"
)

// TemplateData contains data passed to templates
type TemplateData struct {
	ProjectName  string
	Architecture string
	License      string
	Year         int

	// Computed fields
	ProjectTitle    string
	ModuleName      string
	MainPackagePath string

	// Configuration
	WithDocker   bool
	WithMakefile bool
	WithGit      bool

	// Architecture specific data
	ArchData interface{}
}

// NewTemplateData creates template data from config
func NewTemplateData(config *Config) *TemplateData {
	data := &TemplateData{
		ProjectName:  config.ProjectName,
		Architecture: config.Architecture,
		License:      config.License,
		Year:         time.Now().Year(),
		WithDocker:   config.WithDocker,
		WithMakefile: config.WithMakefile,
		WithGit:      config.WithGit,
	}

	// Computed fields
	data.ProjectTitle = strings.Title(config.ProjectName)
	data.ModuleName = config.ProjectName
	data.MainPackagePath = fmt.Sprintf("cmd/%s", config.ProjectName)

	// Set architecture specific data
	switch config.Architecture {
	case "hexagonal":
		data.ArchData = NewHexagonalData(config)
	case "clean":
		data.ArchData = NewCleanData(config)
	case "mvc":
		data.ArchData = NewMVCData(config)
	case "basic":
		data.ArchData = NewBasicData(config)
	}

	return data
}

// Architecture specific data structures

type HexagonalData struct {
	CorePorts []string
	Adapters  []string
	Services  []string
}

func NewHexagonalData(config *Config) *HexagonalData {
	return &HexagonalData{
		CorePorts: []string{"Repository", "Cache", "Logger"},
		Adapters:  []string{"HTTP", "Database", "Cache"},
		Services:  []string{"AppService"},
	}
}

type CleanData struct {
	Entities []string
	UseCases []string
	Handlers []string
}

func NewCleanData(config *Config) *CleanData {
	return &CleanData{
		Entities: []string{"Entity"},
		UseCases: []string{"EntityUseCase"},
		Handlers: []string{"EntityHandler"},
	}
}

type MVCData struct {
	Controllers []string
	Models      []string
	Views       []string
}

func NewMVCData(config *Config) *MVCData {
	return &MVCData{
		Controllers: []string{"BaseController", "HealthController"},
		Models:      []string{"BaseModel"},
		Views:       []string{"Response"},
	}
}

type BasicData struct {
	Handlers []string
	Services []string
}

func NewBasicData(config *Config) *BasicData {
	return &BasicData{
		Handlers: []string{"HealthHandler"},
		Services: []string{"Service"},
	}
}
