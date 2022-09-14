package function

/*
This function template responds with a list of environment variables that starts with TEST_
*/

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	res.Header().Add("Content-Type", "text/plain")
	testEnvVars := []string{}
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "TEST_") {
			testEnvVars = append(testEnvVars, e)
		}
	}
	_, err := fmt.Fprintf(res, "Envs Vars starting with TEST_\n%v\n", strings.Join(testEnvVars, "\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error or response write: %v", err)
	}

}
