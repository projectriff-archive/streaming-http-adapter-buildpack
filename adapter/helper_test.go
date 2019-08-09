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
	layers2 "github.com/buildpack/libbuildpack/layers"
	"github.com/cloudfoundry/libcfbuildpack/layers"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"testing"
)

func TestHelper(t *testing.T) {
	spec.Run(t, "Helper", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		it("rewrites commands", func() {
			m := layers.Metadata{
				Processes: []layers2.Process{
					{Command: "foo bar"},
				},
			}

			Adapt(m)

			g.Expect(m.Processes[0].Command).To(Equal(fmt.Sprintf("%s foo bar", executableName)))
		})
	}, spec.Report(report.Terminal{}))
}
