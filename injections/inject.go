package injections

import "context"

func (i *Injections) Inject(ctx context.Context) (map[string]string, error) {

	return make(map[string]string), nil
}
