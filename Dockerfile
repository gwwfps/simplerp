FROM scratch
EXPOSE 8080
COPY simplerp /
ENTRYPOINT ["/simplerp"]
