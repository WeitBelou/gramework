// Copyright 2017 Kirill Danshin and Gramework contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//

package client

import (
	"errors"

	"github.com/valyala/bytebufferpool"
)

var (
	// ErrNoServerAvailable occurred when no server available in the pool
	ErrNoServerAvailable = errors.New("no server available")

	buffer bytebufferpool.Pool
)
