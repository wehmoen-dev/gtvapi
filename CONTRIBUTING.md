# Contributing to gtvapi

Thank you for your interest in contributing to gtvapi! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and constructive in all interactions. We aim to maintain a welcoming and inclusive community.

## How to Contribute

### Reporting Bugs

Before submitting a bug report:
1. Check the existing issues to avoid duplicates
2. Collect information about your environment (Go version, OS, etc.)
3. Provide a minimal reproducible example if possible

Create an issue with:
- A clear, descriptive title
- Steps to reproduce the problem
- Expected behavior
- Actual behavior
- Any relevant logs or error messages

### Suggesting Features

Feature suggestions are welcome! Please:
1. Check existing issues for similar proposals
2. Clearly describe the feature and its use case
3. Explain why it would be beneficial
4. Consider providing a rough implementation outline

### Pull Requests

1. **Fork the repository** and create a feature branch
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following the coding standards below

3. **Add tests** for your changes
   - Unit tests for new functionality
   - Update existing tests if needed
   - Ensure all tests pass: `go test ./...`

4. **Update documentation**
   - Add GoDoc comments for new public APIs
   - Update README.md if needed
   - Add entries to CHANGELOG.md

5. **Run quality checks**
   ```bash
   make lint
   make test
   make fmt
   ```

6. **Commit your changes** with clear, descriptive messages
   ```bash
   git commit -m "Add feature: description of your changes"
   ```

7. **Push to your fork** and submit a pull request

8. **Respond to feedback** during the review process

## Coding Standards

### Go Style

- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` or `gofumpt` for formatting
- Follow idiomatic Go patterns
- Keep functions small and focused

### Documentation

- Add GoDoc comments for all exported types, functions, and constants
- Use complete sentences in comments
- Provide examples for complex functionality
- Keep documentation up-to-date with code changes

### Testing

- Write table-driven tests where appropriate
- Use meaningful test names that describe what is being tested
- Test edge cases and error conditions
- Aim for high test coverage (target: 80%+)
- Use mocks for external dependencies

### Error Handling

- Return errors rather than panicking
- Wrap errors with context using `fmt.Errorf` with `%w`
- Use meaningful error messages
- Check all errors

### Naming Conventions

- Use descriptive names for variables and functions
- Follow Go naming conventions (camelCase for unexported, PascalCase for exported)
- Avoid abbreviations unless they're widely understood
- Use consistent terminology throughout the codebase

## Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/wehmoen-dev/gtvapi.git
   cd gtvapi
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run tests**
   ```bash
   go test -v ./...
   ```

4. **Run linting**
   ```bash
   golangci-lint run
   ```

## Commit Message Guidelines

Use clear and descriptive commit messages:

- **feat:** A new feature
- **fix:** A bug fix
- **docs:** Documentation changes
- **style:** Code style changes (formatting, missing semicolons, etc.)
- **refactor:** Code refactoring without changing functionality
- **test:** Adding or updating tests
- **chore:** Maintenance tasks (dependency updates, build changes, etc.)

Example:
```
feat: add retry logic with exponential backoff

- Implement retry mechanism for failed API requests
- Add configurable retry parameters in ClientConfig
- Update tests to cover retry scenarios
```

## Questions?

If you have questions about contributing, feel free to open an issue for discussion.

Thank you for contributing to gtvapi!
