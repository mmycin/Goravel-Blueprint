#!/usr/bin/env bash
set -e

APP_NAME="goraveltpl"
ENTRY_POINT="cmd/goraveltpl/main.go"
DIST_DIR="dist"

# Clean previous builds
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

# OS/ARCH combinations
OS_ARCHS=(
  "windows amd64"
  "windows arm64"
  "linux amd64"
  "linux arm64"
  "darwin amd64"
  "darwin arm64"
)

# Build loop
for entry in "${OS_ARCHS[@]}"; do
  read -r GOOS GOARCH <<<"$entry"
  echo "Building $APP_NAME for $GOOS/$GOARCH..."

  OUTPUT_NAME="$APP_NAME-$GOOS-$GOARCH"
  if [ "$GOOS" = "windows" ]; then
    OUTPUT_NAME="$OUTPUT_NAME.exe"
  fi

  # Build
  GOOS=$GOOS GOARCH=$GOARCH go build -o "$DIST_DIR/$OUTPUT_NAME" "$ENTRY_POINT"

done

echo "All binaries are built and packaged in $DIST_DIR/"
