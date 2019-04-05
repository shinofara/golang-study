FROM golang:1.12

RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]

# docker image build --tag mfkessai/codetest:latest .
# docker container run -v (pwd):/work -w /work -p 8888:8888 mfkessai/codetest:latest
