<p align="center">
  <img src="https://pub-fbb2435be97c492d8ece0578844483ea.r2.dev/scailo-logo.png" alt="Scailo Logo" height="80"/>
</p>

<h1 align="center">Scailo Bootstrap CLI</h1>

<p align="center">
  A powerful command-line tool to simplify and accelerate data onboarding into <a target="_blank" href="https://scailo.com">Scailo</a>.
</p>

---

## Overview

`scailo-bootstrap` is a command-line utility built in Go that helps users bootstrap their Scailo instance by uploading all relevant data through CSV files. This removes the need for manual data entry through the Scailo UI and allows for bulk operations in a safe, repeatable manner.

Whether you're setting up a new environment or updating existing records, this tool ensures a smooth, automated process.

---

## Features

- 🔄 **Two Operating Modes**: `download` and `upload`
- 📂 **Bulk Upload via CSV**: Upload complete datasets in one go
- 📥 **CSV Template Generator**: Automatically download CSV templates required for upload
- 🔐 **Secure Authentication**: Environment-based credential management
- ✅ **Idempotent Uploads**: Safe to re-upload files—no duplicate records
- ⚙️ **Custom Form Field Support**: Create and use dynamic form fields seamlessly

---

## Installation

To install the CLI tool, run:

```bash
go install github.com/scailo/scailo-bootstrap@latest
````

This will make the `scailo-bootstrap` command available in your terminal.

---

## Authentication

Before running the tool, you must configure your Scailo credentials. These are stored in a `.env` file located in the root of your working directory.

### Sample `.env` File

```env
server_url=your-scailo-instance.com:443
username=your-username
password=your-password
```

You can also override `server_url` via command-line argument.

---

## Usage

### Basic Command Format

```bash
scailo-bootstrap --mode=<download|upload> --dir=<directory-path> [options]
```

### Available Flags

| Flag               | Description                                                                                      |
| ------------------ | ------------------------------------------------------------------------------------------------ |
| `--mode`           | **(Required)** Operation mode: `download` or `upload`                                            |
| `--dir`            | **(Required)** Directory to read from or write to                                                |
| `--server_url`     | Optional. Overrides `server_url` from `.env`                                                     |
| `--env-file`       | Optional. Path to `.env` file (default: `.env`)                                                  |
| `--should-verify`  | (Upload only) Whether uploaded records should be verified (`true`/`false`). Defaults to `false`. |
| `--should-approve` | (Upload only) Whether uploaded records should be approved (`true`/`false`). Defaults to `false`. |

---

## Examples

### 1. Download CSV Templates

```bash
scailo-bootstrap --mode=download --dir=templates
```

This will download all required CSV templates into the `templates/` folder.

---

### 2. Upload CSV Data

```bash
scailo-bootstrap --mode=upload --dir=uploads --should-verify=true --should-approve=true
```

This uploads all supported CSV files from the `uploads/` directory, verifying and approving the records as specified.

---

### 3. Upload with Custom `.env` File

```bash
scailo-bootstrap --mode=upload --dir=uploads --should-verify=true --should-approve=true --env-file=.env
```

Specifies a custom `.env` file for credentials and server URL.

---

## Custom Form Field Workflow

Scailo allows users to define custom form sections and fields. Here's how to incorporate them into your data upload process:

### Step-by-Step Process

1. **Download Templates**

   ```bash
   scailo-bootstrap --mode=download --dir=templates
   ```

2. **Setup Form Fields**
   Copy the following files from the `templates/` folder to a new directory:

   * `formsections.csv`
   * `formfields.csv`

3. **Fill and Upload Form Fields**
   Populate the above CSVs with your custom fields, then run:

   ```bash
   scailo-bootstrap --mode=upload --dir=form-definitions
   ```

4. **Re-Download Templates**
   Download the templates again to refresh them with your custom field definitions.

   ```bash
   scailo-bootstrap --mode=download --dir=updated-templates
   ```

5. **Populate and Upload Final Data**
   Copy the relevant files from `updated-templates/` into a new folder, fill in your data, and run the upload again:

   ```bash
   scailo-bootstrap --mode=upload --dir=final-data --should-verify=true --should-approve=true
   ```

---

## Idempotency Guarantee

Uploads via this tool are **idempotent**. If a record with the same `CODE` already exists, it is **safely skipped** during upload. This means:

* You can upload the same file multiple times without causing duplication
* It is safe to re-run the upload command at any point

---

## Contributing

This is an internal Scailo CLI tool. If you'd like to contribute or request features, please contact the Scailo development team via [https://scailo.com](https://scailo.com).

---

## Learn More

To understand how Scailo works and how it can help your business, visit the [Scailo Website](https://scailo.com).