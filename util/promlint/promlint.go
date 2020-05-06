// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package promlint provides a linter for Prometheus metrics.
//
// Deprecated: Use the "github.com/prometheus/client_golang/prometheus/testutil/promlint" package instead.
package promlint

import (
	"io"

	"github.com/prometheus/client_golang/prometheus/testutil/promlint"
)

// A Linter is a Prometheus metrics linter.  It identifies issues with metric
// names, types, and metadata, and reports them to the caller.
type Linter struct {
	r io.Reader
}

// A Problem is an issue detected by a Linter.
type Problem struct {
	// The name of the metric indicated by this Problem.
	Metric string

	// A description of the issue for this Problem.
	Text string
}

// New creates a new Linter that reads an input stream of Prometheus metrics.
// Only the text exposition format is supported.
func New(r io.Reader) *Linter {
	return &Linter{
		r: r,
	}
}

// Lint performs a linting pass, returning a slice of Problems indicating any
// issues found in the metrics stream.  The slice is sorted by metric name
// and issue description.
func (l *Linter) Lint() ([]Problem, error) {
	linter := promlint.New(l.r)
	cProblems, err := linter.Lint()
	if err != nil {
		return nil, err
	}
	if len(cProblems) == 0 {
		return nil, nil
	}

	problems := make([]Problem, 0, len(cProblems))
	for i := range cProblems {
		problems = append(problems, Problem{Metric: cProblems[i].Metric, Text: cProblems[i].Text})
	}

	return problems, err
}
