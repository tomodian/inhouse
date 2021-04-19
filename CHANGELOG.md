# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Defined CLIENV for cli/source/test switching

### Changed

- Moved cli to cli/inhouse for go get install
- Renamed module to github.com/tomodian/inhouse

## [0.2.0] - 2021-04-19

### Added

- Added SourcesContainsPWD and TestContainsPWD for simpler test integrations
- Ignored examples directory from code coverage tracking
- Added examples to demonstrate inhouse to be used in arbitary test codes
- Extended Code struct with decorators (colon, csv, tsv, json)
- Added Mozilla Public License 2.0

### Changed

- Prefixed log with DEBUG marker
- Output coverage report after test

### Fixed

- Incorrect path was returned when tilda is specified

## [0.1.0] - 2021-04-18

### Added

- Installed github.com/mitchellh/gox
- Installed github.com/urfave/cli
- Put CHANGELOGs for each modules and directories for narrower change management
- Added Combine to merge all codes
- Log message to STDOUT when DEBUG environment variable is not empty
- Return scan results with line numbers instead of just filepaths
- Added README
- Added DirTestContains checker
- Added Contains checker
- Added test utility for DRYing
- Added AST parser for exported/private functions
- Added exported functions to lookup implementation and test files
- Added globber
- GitHub Actions for Linux testing

### Changed

- Fallback to directory when user specified a filepath as directory parameter
- Enabled to inject preferred directory for globbing
- DRYed checkers
- Rebranded from clubrule to inhouse
- Renamed Dir prefixes for simplicity
