// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package converter

import (
	"context"
	"strings"

	"github.com/drone/drone/core"
	"github.com/drone/drone/plugin/converter/starlark"
)

// New returns a conversion service that converts the
// starlark file to a yaml file.
func New(enabled bool) core.ConvertService {
	return &starlarkPlugin{
		enabled: enabled,
	}
}

type starlarkPlugin struct {
	enabled bool
}

func (p *starlarkPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}

	// if the file extension is not jsonnet we can
	// skip this plugin by returning zero values.
	switch {
	case strings.HasSuffix(req.Repo.Config, ".script"):
	case strings.HasSuffix(req.Repo.Config, ".star"):
	case strings.HasSuffix(req.Repo.Config, ".starlark"):
	default:
		return nil, nil
	}

	file, _ := starlark.Parse(req, nil, nil)
	return &core.Config{
		Data: *file,
	}, nil
}
