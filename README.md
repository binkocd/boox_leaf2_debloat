# Boox Leaf 2 Debloater

A Go-based utility to safely debloat the Onyx Boox Leaf 2. This tool prioritizes privacy by disabling telemetry and background services while preserving core E-Ink functionality.

## 🚀 Features
* **Auto-Backup:** Saves a timestamped CSV of all packages before changes.
* **Safe Mode:** Uses `pm disable-user`, making all changes reversible.
* **Full Revert:** Restore your device state using the generated backup file.
* **Cross-Platform:** Binaries available for Windows, macOS, and Linux.

## 🛠 Prerequisites
1.  **ADB:** Must be installed and in your system `PATH`.
2.  **USB Debugging:** Enabled in `Settings > Developer Options`.
3.  **Go (Optional):** Only needed if building from source (v1.21+).

## 📦 Usage

### Using Pre-built Binaries:
1. Download the latest release for your OS from the **Releases** tab.
2. Open a terminal in that folder.
3. Run:
   ```bash
   ./leaf2-debloat
   ```

### To Revert (Undo everything):
```./leaf2-debloat --revert <your_backup_file>.csv```

## 📋 Targeted Packages
Package	Description
onyx.aiassistant	Telemetry & AI Services
onyx.appmarket	Boox Store
onyx.ksync	Cloud Syncing
onyx.igetshop	Shopping/Ads
onyx.mail	Stock Email Client

## 🤝 Contributing

Contributions are welcome! If you'd like to add a package to the debloat list or improve the tool:

    Fork the repository.

    Create a Feature Branch (git checkout -b feature/AmazingFeature).

    Commit your changes (git commit -m 'Add some AmazingFeature').

    Push to the branch (git push origin feature/AmazingFeature).

    Open a Pull Request.

## 🛠 Development

To run linting and ensure code quality:

## Run golangci-lint
golangci-lint run

## Run tests
go test ./...

## ⚖️ Disclaimer

Use at your own risk. While disable-user is generally safe, disabling system services can cause unexpected behavior.