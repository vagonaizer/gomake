package generator

// ProjectStructure defines the directory structure of a project
type ProjectStructure struct {
	Directories []string
	Files       map[string]string // path -> content
}

// NewProjectStructure creates a new project structure
func NewProjectStructure() *ProjectStructure {
	return &ProjectStructure{
		Directories: make([]string, 0),
		Files:       make(map[string]string),
	}
}

// AddDirectory adds a directory to the structure
func (ps *ProjectStructure) AddDirectory(path string) {
	ps.Directories = append(ps.Directories, path)
}

// AddFile adds a file with content to the structure
func (ps *ProjectStructure) AddFile(path, content string) {
	ps.Files[path] = content
}

// GetDirectories returns all directories
func (ps *ProjectStructure) GetDirectories() []string {
	return ps.Directories
}

// GetFiles returns all files
func (ps *ProjectStructure) GetFiles() map[string]string {
	return ps.Files
}
