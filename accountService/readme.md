To install dependencies locally:

    $ go get -d ./...

You can then build and run the Docker image:

    $ docker build -t lookup .
    $ docker run -p 8080:8080 -it --rm --name accountService accountService

Note: go-wrapper run includes set -x so the binary name is printed to stderr on application startup. If this behavior is undesirable, then switching to CMD ["app"] (or CMD ["myapp"] if a Go custom import path is in use) will silence it by running the built binary directly.

URL to run from browser:
http://localhost:8080/accountService/<name>