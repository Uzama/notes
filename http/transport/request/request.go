package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Decode(ctx context.Context, r *http.Request, decoder DecoderInterface) error {

	bite, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(bite, decoder)
	if err != nil {
		return nil
	}

	return nil
}
