#!/usr/bin/env bash
if ! type go > /dev/null; then
  echo "you need to install golang to continue" && exit 1
fi

TEMP_DIR=$(mktemp -d)
echo "Cloning repo..."

git clone https://git.gay/gb/gb "$TEMP_DIR"

cd "$TEMP_DIR" || (echo "Failed to clone into repo" && exit)

echo "Installing..."
go install -ldflags="-w -s" -gcflags=all=-l .

GOBIN="$(go env GOBIN | awk -F':' '{print $1}')"
GOPATH="$(go env GOPATH | awk -F':' '{print $1}')"/bin
if [ -z "$GOBIN" ]; then
  GOPATH_BIN=$GOPATH
else
  GOPATH_BIN=$GOBIN
fi


echo "Done!"
echo "You can now use $GOPATH_BIN/gb to use the cli"
echo "To use this cli globally, add the following to \$PATH:"
echo "$GOPATH_BIN"

echo "Cleaning up..."
rm -rf "$TEMP_DIR"