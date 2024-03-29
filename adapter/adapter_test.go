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
	"path/filepath"
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestAdapter(t *testing.T) {
	spec.Run(t, "Adapter", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var f *test.BuildFactory

		it.Before(func() {
			f = test.NewBuildFactory(t)
		})

		it("returns true if build plan exists", func() {
			f.AddDependency(adapterExecutableId, filepath.Join("testdata", "streaming-adapter.tgz"))
			f.AddBuildPlan(Dependency, buildplan.Dependency{})

			_, ok, err := NewAdapter(f.Build)
			g.Expect(ok).To(BeTrue())
			g.Expect(err).NotTo(HaveOccurred())
		})

		it("returns false if build plan does not exist", func() {
			_, ok, err := NewAdapter(f.Build)
			g.Expect(ok).To(BeFalse())
			g.Expect(err).NotTo(HaveOccurred())
		})

		it("contributes adapter executable", func() {
			f.AddDependency(adapterExecutableId, filepath.Join("testdata", "streaming-adapter.tgz"))
			f.AddBuildPlan(Dependency, buildplan.Dependency{})

			a, _, err := NewAdapter(f.Build)
			g.Expect(err).NotTo(HaveOccurred())

			g.Expect(a.Contribute()).To(Succeed())

			layer := f.Build.Layers.Layer("streaming-http-adapter")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(layer).To(test.HaveAppendPathLaunchEnvironment("PATH", layer.Root))
		})
	}, spec.Report(report.Terminal{}))
}
