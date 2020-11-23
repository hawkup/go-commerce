# GO commerce

## Tool
Run this command to install Bazel
`$ brew install bazel`

## Setup
Run command for setup file
`$ bazel run //:gazelle`

## Config use go.mod
`$ bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies`

## Update importpath on BUILD.bazel

## Build
`bazel build //packages/order:order`

## Run
`bazel run //packages/order:order`

### Referrence
* https://github.com/bazelbuild/bazel-gazelle#running-gazelle-with-bazel