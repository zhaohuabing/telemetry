// Copyright (c) Tetrate, Inc 2023.
//
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

package telemetry

import (
	"context"
	"errors"
	"testing"
)

func TestNoopLogger(t *testing.T) {
	tests := []struct {
		name        string
		logfunc     func(Logger)
		metricCount float64
	}{
		{"info-", func(l Logger) { l.Info("text", "where", "there") }, 1},
		{"error", func(l Logger) { l.Error("text", errors.New("error"), "where", "there") }, 1},
		{"debug", func(l Logger) { l.Debug("text", "where", "there") }, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			metric := mockMetric{}
			ctx := KeyValuesToContext(context.Background(), "ctx", "value")
			l := NoopLogger().Context(ctx).Metric(&metric).With().With("lvl", LevelInfo).With("missing")

			tt.logfunc(l)

			if metric.count != 0 {
				t.Fatalf("metric.count=%v, want 0", metric.count)
			}
		})
	}
}

type mockMetric struct {
	Metric
	count float64
}

func (m *mockMetric) RecordContext(_ context.Context, value float64) { m.count += value }
