/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Source directory to backup and archive to destination directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Backup Started...")
		fmt.Println("BackupSrc:", src)
		fmt.Println("BackupDest:", dest)
		if err:= performBackup(src, dest, number); err != nil {
			fmt.Println(os.Stderr, "Error:", err)
		}
		fmt.Println("Backup Completed...")
	},
}

var src, dest string
var number int

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringVarP(&src, "src", "s", "", "Source directory to backup")
	backupCmd.Flags().StringVarP(&dest, "dest", "d", "", "Destination directory for backup")
	//backupCmd.Flags().IntVarP(&number, "number", "n", 7, "Number of backups to keep")

	backupCmd.MarkFlagRequired("src")
	backupCmd.MarkFlagRequired("dest")
}

func performBackup(src, dest string, number int) error {
	//Timestampは下記指定じゃないとフォーマットが認識されない
    timestamp := time.Now().Format("20060102150405")
    tarName := fmt.Sprintf("%s.tgz", timestamp)

    tarPath := filepath.Join(dest, tarName)
    cmd := exec.Command("tar", "-cvzf", tarPath, src)
    if err := cmd.Run(); err != nil {
		fmt.Println(os.Stderr, "Error running tar command:", err)
        return err
    }

    //return removeOldBackups(dest, number)
	return nil
}

func removeOldBackups(dest string, number int) error {
    files, err := filepath.Glob(filepath.Join(dest, "*.tgz"))
    if err != nil {
        return err
    }

    if len(files) <= number {
        return nil
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i] < files[j]
    })

    for _, file := range files[:len(files)-number] {
        if err := os.Remove(file); err != nil {
            return err
        }
    }

    return nil
}