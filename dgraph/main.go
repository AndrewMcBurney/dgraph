/*
 * Copyright (C) 2017 Dgraph Labs, Inc. and Contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"math/rand"
	"runtime"
	"time"

	ddtrace "github.com/DataDog/dd-trace-go/opentracing"
	"github.com/dgraph-io/dgraph/dgraph/cmd"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	// create a Tracer configuration
	config := ddtrace.NewConfiguration()
	config.ServiceName = "dgraph"

	// initialize a Tracer and ensure a graceful shutdown
	// using the `closer.Close()`
	tracer, closer, err := ddtrace.NewTracer(config)
	if err != nil {
		// handle the configuration error
	}
	defer closer.Close()

	// set the Datadog tracer as a GlobalTracer
	opentracing.SetGlobalTracer(tracer)

	rand.Seed(time.Now().UnixNano())
	// Setting a higher number here allows more disk I/O calls to be scheduled, hence considerably
	// improving throughput. The extra CPU overhead is almost negligible in comparison. The
	// benchmark notes are located in badger-bench/randread.
	runtime.GOMAXPROCS(128)
	cmd.Execute()
}
