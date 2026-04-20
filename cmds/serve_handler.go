package cmds

import (
	"context"

	"ksogit.kingsoft.net/o/wps"
)

func serveHandler(ctx context.Context, args []string, opts *serveOptions) error {
	wps.Go(AppName, "Version: "+Version+"Git Hash: "+GitHash+"; Build Time: "+BuildTime, opts.flagSet)
	return nil
}
