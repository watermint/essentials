# Essential libraries for Go

# Objectives

## Goals

* All native Go library.
* Minimize `if err != nil` idiom. Compare to `nil` sometime cause unintended bugs.
* Minimize use of external libraries.
* Separate interface from implementation.
* Ease of development.

## Core usecase

This library is intended to use a desktop or command-line applications, not on a server or IoT/mobile.
The library may be optimized for that.
For example, a target memory footprint may not exceed 1-2 GiB for a single instance.

## Non-goals

* Performance (this is often trade-off to ease-of-development)
* Minimal footprints
* Compatibility (because this project is experimental)

# Structure

## Naming

* All submodules should start with the prefix `e`. For example, if a sub-module is for logging, the name should be `elog`.
