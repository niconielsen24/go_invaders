#!/usr/bin/env bash
set -euo pipefail

MISSING=()

check() {
    command -v "$1" &>/dev/null
}

need() {
    echo "  missing: $1"
    MISSING+=("$1")
}

echo "Checking dependencies..."

# Go
if check go; then
    echo "  go $(go version | awk '{print $3}')"
else
    need go
fi

# C compiler (required by CGo, used by raylib-go)
if check gcc; then
    echo "  gcc $(gcc --version | head -1 | awk '{print $3}')"
else
    need gcc
fi

# System libs required by raylib on Linux
LIBS=(
    libgl-dev
    libx11-dev
    libxrandr-dev
    libxinerama-dev
    libxcursor-dev
    libxi-dev
    libxxf86vm-dev
)

for lib in "${LIBS[@]}"; do
    if dpkg -s "$lib" &>/dev/null 2>&1; then
        echo "  $lib"
    else
        need "$lib"
    fi
done

if [ ${#MISSING[@]} -eq 0 ]; then
    echo ""
    echo "All dependencies satisfied."
    exit 0
fi

echo ""
echo "Installing missing dependencies: ${MISSING[*]}"

# Go must be installed manually
if [[ " ${MISSING[*]} " == *" go "* ]]; then
    echo ""
    echo "Go must be installed manually: https://go.dev/dl/"
    MISSING=("${MISSING[@]/go}")
fi

# Install system packages
PKGS=("${MISSING[@]}")
if [ ${#PKGS[@]} -gt 0 ]; then
    if check apt-get; then
        sudo apt-get update -qq && sudo apt-get install -y "${PKGS[@]}"
    elif check dnf; then
        sudo dnf install -y "${PKGS[@]}"
    elif check pacman; then
        sudo pacman -S --noconfirm "${PKGS[@]}"
    else
        echo "Unsupported package manager. Install manually: ${PKGS[*]}"
        exit 1
    fi
fi

echo ""
echo "Done. Run 'cd cmd && go run .' to start the game."
