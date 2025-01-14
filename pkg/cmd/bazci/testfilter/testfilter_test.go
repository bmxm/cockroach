// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package testfilter

import (
	"strings"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/datadriven"
)

func TestFilterAndWrite(t *testing.T) {
	defer leaktest.AfterTest(t)()
	datadriven.Walk(t, testutils.TestDataPath(t), func(t *testing.T, path string) {
		datadriven.RunTest(t, path, func(t *testing.T, td *datadriven.TestData) string {
			in := strings.NewReader(td.Input)
			var out strings.Builder
			if err := FilterAndWrite(in, &out, []string{td.Cmd}); err != nil {
				return err.Error()
			}
			// At the time of writing, datadriven garbles the test files when
			// rewriting a "\n" output, so make sure we never have trailing
			// newlines.
			return strings.TrimRight(out.String(), "\r\n")
		})
	})
}
