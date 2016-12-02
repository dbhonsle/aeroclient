// Copyright 2013-2016 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

// QueryPolicy encapsulates parameters for policy attributes used in query operations.
type QueryPolicy struct {
	*MultiPolicy
}

// NewQueryPolicy generates a new QueryPolicy instance with default values.
func NewQueryPolicy() *QueryPolicy {
	res := &QueryPolicy{
		MultiPolicy: NewMultiPolicy(),
	}

	// Retry policy must be one-shot for queries
	res.MaxRetries = 0

	return res
}
