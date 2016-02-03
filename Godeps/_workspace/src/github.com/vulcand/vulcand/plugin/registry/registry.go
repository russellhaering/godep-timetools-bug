// This file will be generated to include all customer specific middlewares
package registry

import (
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin"
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin/cbreaker"
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin/connlimit"
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin/ratelimit"
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin/rewrite"
	"github.com/russellhaering/godep-timetools-bug/Godeps/_workspace/src/github.com/vulcand/vulcand/plugin/trace"
)

func GetRegistry() *plugin.Registry {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{
		ratelimit.GetSpec(),
		connlimit.GetSpec(),
		rewrite.GetSpec(),
		cbreaker.GetSpec(),
		trace.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			panic(err)
		}
	}

	return r
}
