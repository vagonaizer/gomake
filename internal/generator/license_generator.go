package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LicenseGenerator handles license file generation
type LicenseGenerator struct {
	config *Config
	logger Logger
}

// NewLicenseGenerator creates a new license generator
func NewLicenseGenerator(config *Config, logger Logger) *LicenseGenerator {
	return &LicenseGenerator{
		config: config,
		logger: logger,
	}
}

// Generate creates a license file
func (lg *LicenseGenerator) Generate(projectPath string) error {
	lg.logger.Info("Generating license file", "license", lg.config.License)

	var content string
	year := time.Now().Year()

	switch lg.config.License {
	case "MIT":
		content = lg.generateMITLicense(year)
	case "Apache":
		content = lg.generateApacheLicense(year)
	case "BSD":
		content = lg.generateBSDLicense(year)
	case "GPL":
		content = lg.generateGPLLicense(year)
	default:
		return fmt.Errorf("unsupported license: %s", lg.config.License)
	}

	filePath := filepath.Join(projectPath, "LICENSE")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func (lg *LicenseGenerator) generateMITLicense(year int) string {
	return fmt.Sprintf(`MIT License

Copyright (c) %d %s

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`, year, lg.config.ProjectName)
}

func (lg *LicenseGenerator) generateApacheLicense(year int) string {
	return fmt.Sprintf(`Apache License
Version 2.0, January 2004
http://www.apache.org/licenses/

Copyright %d %s

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
`, year, lg.config.ProjectName)
}

func (lg *LicenseGenerator) generateBSDLicense(year int) string {
	return fmt.Sprintf(`BSD 3-Clause License

Copyright (c) %d, %s
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`, year, lg.config.ProjectName)
}

func (lg *LicenseGenerator) generateGPLLicense(year int) string {
	return fmt.Sprintf(`GNU GENERAL PUBLIC LICENSE
Version 3, 29 June 2007

Copyright (C) %d %s

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
`, year, lg.config.ProjectName)
}
