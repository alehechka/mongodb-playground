# mongodb-playground

Multi-arch Docker builds:

```bash
docker buildx build -t ghcr.io/alehechka/mongodb-playground:latest --platform=linux/arm64,linux/amd64 . --push
```
