#!/usr/bin/env bash

# Build and run initially
go build -o build/lazyrest || exit 1
clear
build/lazyrest -w ./demo/ -d -l debug &
PID=$!

# Watch for changes and rebuild
fswatch -e ".*" -i "\\.go$" -r . | while read -r file; do
    # Build new version
    if go build -o build/lazyrest; then
        # Kill old process gracefully
        kill $PID 2>/dev/null
        sleep 0.2
        
        # Start new version
        clear
        build/lazyrest -w ./demo/ -d -l debug &
        PID=$!
    else
        echo "Build failed, keeping old version running"
    fi
done
