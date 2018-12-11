FROM quay.io/bitriseio/bitrise-base

# envs
ENV PROJ_NAME=bitrise-welcome
ENV BITRISE_SOURCE_DIR="/bitrise/go/src/github.com/bitrise-team/$PROJ_NAME"

# Get go tools
RUN go get github.com/codegangsta/gin \
    && go get github.com/kisielk/errcheck \
    && go get golang.org/x/lint/golint

WORKDIR $BITRISE_SOURCE_DIR

CMD $PROJ_NAME
