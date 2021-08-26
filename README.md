# Essential libraries for Go

# Objectives

## Goals

* All native Go library.
* Minimize use of external libraries.
* Separate interface from implementation.
* Safer code.
* Ease of development.

## Core usecase

This library is intended to use a desktop or command-line applications, not on a server or IoT/mobile. The library may
be optimized for that. For example, a target memory footprint may not exceed 1-2 GiB for a single instance.

## Non-goals

* Performance (this is often trade-off to ease-of-development)
* Minimal footprints
* Compatibility (because this project is experimental)

## Target platform

* Windows 10 or above (x64)
* Darwin (macOS) (x64, arm64)
* Linux (x64)

# Structure

## Naming

* All submodules should start with the prefix `e`. For example, if a sub-module is for logging, the name should
  be `elog`.
