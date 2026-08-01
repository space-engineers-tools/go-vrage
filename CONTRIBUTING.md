# Contribution

Hi there! <img src="https://raw.githubusercontent.com/EchterTimo/EchterTimo/refs/heads/main/images/wave.gif" width="20px" height="20px">

Thank you for your interest in contributing to this project!
We appreciate all contributions, whether it is fixing a typo, improving documentation, or adding a new feature.

Since this project is a library, it is important to maintain a high standard of code quality, documentation, testing and backward compatibility.

To ensure quality, all contributions should comply with the following guidelines.

> [!NOTE]  
> If you differ from any of the guidelines below, please provide a justification in the PR description and it might be merged anyway.

## Guidelines

### 1. Domain Rules and Project Scope

- This project is a client for a HTTP API. It is not a game mod or a server plugin.
- It should focus on providing a stateless and typed interface to the API. Caching or event handling should be done in separate projects.
- Please minimize the number of external  dependencies.
- Since this project is a lib and doesn't need an entry point, the following files should not be included in the repo:
  - `main.go`
  - `cmd/`

### 2. Code Quality

- Please follow the Go coding style and conventions.

### 3. Documentation

Document all functions, methods, types, and struct attributes using [GoDoc comments](https://go.dev/doc/comment).

**A good comment**:

- is as short as possible + as long as necessary
- is written in complete sentences
- explains the "why", not just the "what"
- provides usage examples where appropriate

### 4. Testing

- Add tests for new features and bug fixes.
- Tests should also cover all edge cases.
- Please use [testify](https://github.com/stretchr/testify) for all tests.
- For documentation changes, you can skip adding tests.
- Ensure that all tests pass before submitting a PR.

## Tools

The following tools are recommended for development:

- [golangci-lint](https://github.com/golangci/golangci-lint) is a linter for Go. You can manually run it or register it as a pre-push hook like shown in [.githooks/README.md](.githooks/README.md).
- [pkgsite](https://github.com/golang/pkgsite) is a local version of [pkg.go.dev](https://pkg.go.dev/). It allows you to preview the documentation without publishing it.
