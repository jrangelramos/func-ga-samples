package function

/*
This function template returns on body the content of a file on the server.
This is useful to inspect secrets and config maps
*/

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"strings"
)

func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/plain")

	// v=/test/volume-config/myconfig
	q := strings.Split(req.URL.RawQuery, "=")
	action := q[0]
	path := q[1]

	if action == "v" {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v", err)
		}
		_, err = fmt.Fprintf(res, "Content of file %v\n%v", path, string(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on response write: %v", err)
		}
	}

}