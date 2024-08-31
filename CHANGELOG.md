# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Added additional test cases for various language inputs.
- Improved error handling for empty inputs and unexpected characters.

## [v1.0.1] - 2024-08-29

### Fixed

- Corrected a bug in the regular expression for removing special characters, ensuring better support for UTF-8 characters, including Persian.

## [v1.0.0] - 2024-08-28

### Added

- Initial release of `slugify`.
- Implemented the main `Slugify` function to convert strings into a URL-friendly slug format.
- Supported UTF-8 characters, including Persian, English, and mixed-language input.
- Basic error handling for malformed input.
- Added a comprehensive test suite with 10 different test cases.
