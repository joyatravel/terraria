FROM alpine:3.9

# Copy built binary.
ENV PROGRAM=server
COPY ./dist/${PROGRAM} /bin/${PROGRAM}

# Configure env and exposed ports.
ENV GOENV=production
EXPOSE 3000

ENTRYPOINT $PROGRAM
