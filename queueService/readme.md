To install dependencies locally:

    $ go get -d ./...

You can then build and run the Docker image:

    $ docker build -t queueService .
    $ docker run -p 8082:8082 -it --rm --name queueService queueService

Note: go-wrapper run includes set -x so the binary name is printed to stderr on application startup. If this behavior is undesirable, then switching to CMD ["app"] (or CMD ["myapp"] if a Go custom import path is in use) will silence it by running the built binary directly.

URL to run from curl:
curl --request POST --url http://localhost:8082/register/daniel/0123456