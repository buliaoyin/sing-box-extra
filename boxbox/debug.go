//go:build go1.19

package boxbox

import (
	"runtime/debug"

	"github.com/sagernet/sing-box/common/dialer/conntrack"
	"github.com/sagernet/sing-box/option"
)

func applyDebugOptions(options option.DebugOptions) {
	if options.GCPercent != nil {
		debug.SetGCPercent(*options.GCPercent)
	}
	if options.MaxStack != nil {
		debug.SetMaxStack(*options.MaxStack)
	}
	if options.MaxThreads != nil {
		debug.SetMaxThreads(*options.MaxThreads)
	}
	if options.PanicOnFault != nil {
		debug.SetPanicOnFault(*options.PanicOnFault)
	}
	if options.TraceBack != "" {
		debug.SetTraceback(options.TraceBack)
	}
	if options.MemoryLimit != 0 {
		debug.SetMemoryLimit(int64(options.MemoryLimit))
		conntrack.MemoryLimit = int64(options.MemoryLimit)
	}
	if options.OOMKiller != nil {
		conntrack.KillerEnabled = *options.OOMKiller
	}
}
