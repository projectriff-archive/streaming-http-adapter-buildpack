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
	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

// adapterExecutableId is the id of the dependency containing the adapter executable.
const adapterExecutableId = "streaming-http-adapter"

// Dependency indicates that a foreign buildpack requires the streaming adapter.
const Dependency = "http-proxy"

// ProxyAvailable is a build plan key indicating that the adapter is available and can be used by other buildpacks.
const ProxyAvailable = "proxy-available"

// Adapter represents an Adapter contribution by the buildpack.
type Adapter struct {

	// executable is a layer containing the expanded binary for the streaming adapter.
	executable layers.DependencyLayer
}

// Contribute contributes the streaming adapter binary executable to a launch layer.
func (a Adapter) Contribute() error {
	if err := a.executable.Contribute(func(artifact string, layer layers.DependencyLayer) error {
		layer.Logger.Body("Expanding to %s", layer.Root)
		if err := helper.ExtractTarGz(artifact, layer.Root, 0); err != nil {
			return err
		}
		return layer.AppendPathLaunchEnv("PATH", "%s", layer.Root)
	}, layers.Launch) ; err != nil {
		return err
	}

	return nil
}

// NewAdapter creates a new Streaming Adapter instance. OK is true if build plan contains "http-proxy" dependency, otherwise false.
func NewAdapter(build build.Build) (Adapter, bool, error) {
	bp, ok := build.BuildPlan[Dependency]
	if !ok {
		return Adapter{}, false, nil
	}
	deps, err := build.Buildpack.Dependencies()
	if err != nil {
		return Adapter{}, false, err
	}

	dep, err := deps.Best(adapterExecutableId, bp.Version, build.Stack)
	if err != nil {
		return Adapter{}, false, err
	}

	return Adapter{
		executable: build.Layers.DependencyLayer(dep),
	}, true, nil
}
