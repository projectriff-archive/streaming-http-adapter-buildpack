/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package adapter

import (
	"fmt"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

const executableName = "streaming-http-adapter"

// Adapt rewrites launch metadata so that process definitions start the http streaming adapter and delegate to
// the original process definition.
func Adapt(metadata layers.Metadata) layers.Metadata {
	for i, p := range metadata.Processes {
		p.Command = fmt.Sprintf("%s %s", executableName, p.Command)
		metadata.Processes[i] = p
	}

	return metadata
}