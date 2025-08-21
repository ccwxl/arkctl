/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package deploy

import (
	"reflect"
	"testing"
)

func TestParseMavenCommand(t *testing.T) {
	tests := []struct {
		name     string
		mvnCmd   string
		wantCmd  string
		wantArgs []string
	}{
		{
			name:     "empty command",
			mvnCmd:   "",
			wantCmd:  "mvn",
			wantArgs: []string{"clean", "package", "-Dmaven.test.skip=true"},
		},
		{
			name:     "custom command without args",
			mvnCmd:   "mvnw",
			wantCmd:  "mvnw",
			wantArgs: []string{},
		},
		{
			name:     "custom command with args",
			mvnCmd:   "mvn clean install",
			wantCmd:  "mvn",
			wantArgs: []string{"clean", "install"},
		},
		{
			name:     "complex command with args",
			mvnCmd:   "mvnw clean package -DskipTests",
			wantCmd:  "mvnw",
			wantArgs: []string{"clean", "package", "-DskipTests"},
		},
		{
			name:     "command with multiple spaces",
			mvnCmd:   "mvn    clean    package",
			wantCmd:  "mvn",
			wantArgs: []string{"clean", "package"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs := parseMavenCommand(tt.mvnCmd)
			if gotCmd != tt.wantCmd {
				t.Errorf("parseMavenCommand() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("parseMavenCommand() gotArgs = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
