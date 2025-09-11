# Docker Usage

## Using Pre-built Images

Pre-built Docker images are automatically built and published to GitHub Container Registry:

```bash
# Pull the latest image
docker pull ghcr.io/eurusik/rpi_export:main

# Run the container
docker run -d \
  --name rpi_exporter \
  --privileged \
  -p 9110:9110 \
  ghcr.io/eurusik/rpi_export:main
```

## Available Tags

- `main` - Latest build from main branch
- `v*` - Specific version tags (e.g., `v1.0.0`)
- `sha-<commit>` - Specific commit builds

## Docker Compose

```yaml
version: '3.8'
services:
  rpi_exporter:
    image: ghcr.io/eurusik/rpi_export:main
    container_name: rpi_exporter
    privileged: true
    ports:
      - "9110:9110"
    restart: unless-stopped
```

## Building Locally

```bash
# Build for current platform
docker build -t rpi_exporter .

# Build for multiple platforms
docker buildx build \
  --platform linux/arm/v7,linux/arm64 \
  -t rpi_exporter .
```

## CI/CD Pipeline

The project uses GitHub Actions to automatically:

1. **Build** - Compile Go binary for ARM architectures
2. **Test** - Run tests and linting
3. **Package** - Create multi-architecture Docker images
4. **Publish** - Push to GitHub Container Registry

### Triggers

- **Push to main** - Builds and publishes `main` tag
- **Tags** - Builds and publishes version tags
- **Pull Requests** - Builds but doesn't publish

### Supported Architectures

- `linux/arm/v7` - Raspberry Pi 2/3/4 (32-bit)
- `linux/arm64` - Raspberry Pi 3/4 (64-bit)

## Permissions

The container requires `--privileged` mode to access hardware information through `/dev/vcio` device.