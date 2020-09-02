# Changelog
Starting with version 0.3.0, all notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security

## 0.5.1

### Added
- Some extra logging messages

### Fixed
- An empty column heading will terminate column name processing
- Empty rows are ignored
- Empty cells in otherwise valid rows are completely left out of table entries (so they get inserted as nulls or 
  defaults when the database is seeded)
- Setting `LOGLEVEL` environment variable actually affects loglevel

## 0.5.0

### Added
- Generating seed YAML files from an Excel workbook

## 0.4.1

### Fixed
- Creating new Migration file uses MIGRATIONS_DIRECTORY instead of hard-coded value

## 0.4.0

### Added
- Migrations directory now a config value
- Seed directory now a config value
- Loglevel can now be set in the config
- ROOT_DB_NAME now defaults to `postgres`

### Changed
- Refactored files into a more standard project layout

### Removed
- Root configuration

## 0.3.2

### Fixed
- Updated migrations directory from db/migrations to database/migrations

## 0.3.1

### Fixed
- Now reading the loglevel from the correct config key

## 0.3.0

### Changed
- `db` commands were promoted to the root level

### Removed
- Project generators
- Placeholder code for non-implemented commands 
