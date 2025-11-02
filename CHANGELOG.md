# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive README.md with usage examples and API documentation
- LICENSE file (MIT)
- Context support for all API methods
- Retry logic with exponential backoff
- Custom client configuration options
- Improved error handling with detailed error messages
- Full GoDoc documentation for all exported types and functions
- Comprehensive test suite with mocked HTTP responses
- Example tests demonstrating library usage
- Benchmark tests for performance monitoring
- CI/CD configuration with GitHub Actions
- golangci-lint configuration for code quality
- Makefile for common development tasks
- Input validation for all API methods
- Better type safety and consistent return types

### Changed
- Refactored client structure for better maintainability
- Improved method signatures for consistency
- Enhanced type definitions with proper documentation
- Optimized HTTP client configuration

### Fixed
- Nil pointer dereference issues in error handling
- Inconsistent error handling across methods

## [1.0.0] - Initial Release

### Added
- Basic Gronkh.TV API client functionality
- Video information retrieval
- Video comments fetching
- Video playlist URL retrieval
- Video discovery features
- Tag listing
- Twitch live status checking
