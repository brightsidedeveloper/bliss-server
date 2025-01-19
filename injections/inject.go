package injections

import "context"

type ExampleParams struct {
	Wow string
}

func (i *Injections) Inject(ctx context.Context, arg ExampleParams) (map[string]string, error) {

	res := make(map[string]string)

	res["wow"] = arg.Wow

	return res, nil
}
