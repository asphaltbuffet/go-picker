FROM scratch
COPY go-picker /
ENTRYPOINT ["/go-picker"]
