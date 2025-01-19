package injections

import "context"

type InjectParams struct {
	Wow string
}

func (i *Injections) Inject(ctx context.Context, arg InjectParams) (map[string]string, error) {

	res := make(map[string]string)

	res["wow"] = arg.Wow

	return res, nil
}
