package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Bloatware struct {
	ID          string
	Description string
}

func main() {
	// Check for revert flag: go run main.go --revert backup.csv
	if len(os.Args) > 1 && os.Args[1] == "--revert" {
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please provide the backup CSV file path.")
			return
		}
		revertFromBackup(os.Args[2])
		return
	}

	apps := []Bloatware{
		{"com.onyx.aiassistant", "Onyx AI Assistant & Telemetry"},
		{"com.onyx.appmarket", "Onyx App Store"},
		{"com.onyx.android.ksync", "Boox Cloud Sync"},
		{"com.onyx.igetshop", "Shopping/Ads"},
		{"com.onyx.mail", "Stock Mail Client"},
		{"com.onyx.easytransfer", "BooxDrop File Transfer"},
	}

	fmt.Println("🔍 Step 1: Backing up current package list...")
	backupFile, err := backupPackages()
	if err != nil {
		fmt.Printf("❌ Backup failed: %v\n", err)
		return
	}

	fmt.Println("🚀 Step 2: Starting Boox Leaf 2 Debloat...")
	for _, app := range apps {
		fmt.Printf("📦 Disabling: %s...", app.ID)
		cmd := exec.Command("adb", "shell", "pm", "disable-user", "--user", "0", app.ID)
		if err := cmd.Run(); err != nil {
			fmt.Printf(" ❌ Failed (likely already disabled)\n")
		} else {
			fmt.Println(" ✅ Done")
		}
	}
	fmt.Printf("\n✨ Success! To undo this, run:\ngo run main.go --revert %s\n", backupFile)
}

func backupPackages() (string, error) {
	out, err := exec.Command("adb", "shell", "pm", "list", "packages").Output()
	if err != nil {
		return "", err
	}

	timestamp := time.Now().Format("2006-01-02_150405")
	filename := fmt.Sprintf("leaf2_backup_%s.csv", timestamp)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"package_name"})
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		pkg := strings.TrimSpace(strings.TrimPrefix(line, "package:"))
		if pkg != "" {
			writer.Write([]string{pkg})
		}
	}
	return filename, nil
}

func revertFromBackup(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ Could not open backup file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	fmt.Println("🔄 Reverting packages to 'Enabled' state...")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		pkg := record[0]
		fmt.Printf("🔓 Enabling: %s...", pkg)
		exec.Command("adb", "shell", "pm", "enable", pkg).Run()
		fmt.Println(" ✅")
	}
	fmt.Println("✨ Revert complete.")
}